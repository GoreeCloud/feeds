package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	feedsserver "github.com/GoreeCloud/feeds-server/internal/server"
	feedspostgres "github.com/GoreeCloud/feeds-server/internal/storage/postgres"
)

const databaseURLEnvironment = "GOREECLOUD_FEEDS_DATABASE_URL"

func main() {
	listen := flag.String("listen", "127.0.0.1:8080", "development listen address")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	handler, closeRuntime, err := buildHandler(ctx, os.Getenv(databaseURLEnvironment))
	if err != nil {
		log.Fatalf("configure Development runtime: %v", err)
	}
	defer closeRuntime()

	srv := &http.Server{
		Addr:              *listen,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		}
	}()

	log.Printf("GoreeCloud Feeds Server Development runtime listening on %s", *listen)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed: %v", err)
	}
}

func buildHandler(
	ctx context.Context,
	databaseURL string,
) (http.Handler, func(), error) {
	if strings.TrimSpace(databaseURL) == "" {
		return feedsserver.NewHandler(), func() {}, nil
	}

	store, err := feedspostgres.Open(ctx, databaseURL)
	if err != nil {
		return nil, nil, fmt.Errorf("open Development PostgreSQL store: %w", err)
	}
	if err := store.Migrate(ctx); err != nil {
		store.Close()
		return nil, nil, fmt.Errorf("migrate Development PostgreSQL store: %w", err)
	}

	// The article reader may be wired independently from user identity.
	// Without an approved UserContextResolver, the server intentionally does
	// not advertise articles:list-v1 and GET /api/v1/articles remains 503.
	handler := feedsserver.NewHandlerWithDependencies(feedsserver.Dependencies{
		ArticleReader: store,
	})

	return handler, store.Close, nil
}
