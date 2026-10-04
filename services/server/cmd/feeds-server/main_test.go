package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBuildHandlerWithoutDatabaseKeepsArticleListingDisabled(t *testing.T) {
	handler, closeRuntime, err := buildHandler(context.Background(), "")
	if err != nil {
		t.Fatalf("buildHandler() error = %v", err)
	}
	defer closeRuntime()

	assertArticleCapabilityDisabled(t, handler)
}

func TestBuildHandlerWithDevelopmentDatabaseMigratesButKeepsArticleListingDisabled(t *testing.T) {
	databaseURL := os.Getenv("FEEDS_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("FEEDS_TEST_DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	handler, closeRuntime, err := buildHandler(ctx, databaseURL)
	if err != nil {
		t.Fatalf("buildHandler() error = %v", err)
	}
	defer closeRuntime()

	assertArticleCapabilityDisabled(t, handler)
}

func TestBuildHandlerDatabaseFailureDoesNotEchoConnectionSecret(t *testing.T) {
	const secret = "do-not-log-this-password"
	_, _, err := buildHandler(
		context.Background(),
		"postgres://feeds:"+secret+"@127.0.0.1:1/feeds?sslmode=disable",
	)
	if err == nil {
		t.Fatal("expected unreachable database to fail")
	}
	if strings.Contains(err.Error(), secret) {
		t.Fatalf("database error leaked connection secret: %v", err)
	}
}

func assertArticleCapabilityDisabled(t *testing.T, handler http.Handler) {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/capabilities", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("capabilities: expected %d, got %d", http.StatusOK, res.Code)
	}

	var body struct {
		Capabilities []string `json:"capabilities"`
	}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode capabilities: %v", err)
	}
	for _, capability := range body.Capabilities {
		if capability == "articles:list-v1" {
			t.Fatal("article capability must remain disabled without a user-context resolver")
		}
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
	res = httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	if res.Code != http.StatusServiceUnavailable {
		t.Fatalf("articles: expected %d, got %d", http.StatusServiceUnavailable, res.Code)
	}
}
