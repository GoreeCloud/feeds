package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/GoreeCloud/feeds-server/internal/feed"
)

type fakeArticleReader struct {
	articles  []feed.UserArticleSummary
	err       error
	gotUserID feed.ID
	gotLimit  int
	callCount int
}

func (f *fakeArticleReader) ListRecentArticlesForUser(
	_ context.Context,
	userID feed.ID,
	limit int,
) ([]feed.UserArticleSummary, error) {
	f.callCount++
	f.gotUserID = userID
	f.gotLimit = limit
	return f.articles, f.err
}

type fakeUserContextResolver struct {
	userID    feed.ID
	err       error
	callCount int
}

func (f *fakeUserContextResolver) ResolveUserID(_ *http.Request) (feed.ID, error) {
	f.callCount++
	return f.userID, f.err
}

func TestArticleListCapabilityRequiresReaderAndUserContext(t *testing.T) {
	reader := &fakeArticleReader{}
	resolver := &fakeUserContextResolver{userID: "trusted-user"}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/capabilities", nil)
	res := httptest.NewRecorder()
	NewHandlerWithDependencies(Dependencies{
		ArticleReader: reader,
		UserContext:   resolver,
	}).ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, res.Code)
	}
	var body capabilityResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode capability response: %v", err)
	}
	if len(body.Capabilities) != 1 || body.Capabilities[0] != ArticleListCapability {
		t.Fatalf("unexpected capabilities: %v", body.Capabilities)
	}
}

func TestArticleListUsesOnlyServerDerivedUserContext(t *testing.T) {
	published := time.Date(2026, 10, 4, 16, 0, 0, 0, time.UTC)
	reader := &fakeArticleReader{
		articles: []feed.UserArticleSummary{
			{
				ArticleID:   "article-1",
				FeedID:      "feed-1",
				FeedTitle:   "Example Feed",
				URL:         "https://example.test/articles/1",
				Title:       "Article One",
				Author:      "Example Author",
				PublishedAt: &published,
				Summary:     "Summary",
				Language:    "en",
				Read:        true,
				Saved:       true,
			},
		},
	}
	resolver := &fakeUserContextResolver{userID: "trusted-user"}
	handler := NewHandlerWithDependencies(Dependencies{
		ArticleReader: reader,
		UserContext:   resolver,
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles?limit=2", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d: %s", http.StatusOK, res.Code, res.Body.String())
	}
	if resolver.callCount != 1 {
		t.Fatalf("expected resolver once, got %d", resolver.callCount)
	}
	if reader.callCount != 1 || reader.gotUserID != "trusted-user" || reader.gotLimit != 2 {
		t.Fatalf("unexpected reader call: count=%d user=%q limit=%d", reader.callCount, reader.gotUserID, reader.gotLimit)
	}

	var body articleListResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode article response: %v", err)
	}
	if len(body.Articles) != 1 {
		t.Fatalf("expected 1 article, got %d", len(body.Articles))
	}
	if body.Articles[0].ID != "article-1" || body.Articles[0].FeedTitle != "Example Feed" {
		t.Fatalf("unexpected article response: %+v", body.Articles[0])
	}
	if body.Articles[0].PublishedAt == nil || !body.Articles[0].PublishedAt.Equal(published) {
		t.Fatalf("unexpected publication time: %v", body.Articles[0].PublishedAt)
	}
}

func TestArticleListRejectsClientSuppliedUserIDBeforeResolvingContext(t *testing.T) {
	reader := &fakeArticleReader{}
	resolver := &fakeUserContextResolver{userID: "trusted-user"}
	handler := NewHandlerWithDependencies(Dependencies{
		ArticleReader: reader,
		UserContext:   resolver,
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles?user_id=attacker", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
	}
	if resolver.callCount != 0 || reader.callCount != 0 {
		t.Fatalf("client user id must fail before dependency calls: resolver=%d reader=%d", resolver.callCount, reader.callCount)
	}
}

func TestArticleListFailsClosedWithoutConfiguredContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
	res := httptest.NewRecorder()
	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected %d, got %d", http.StatusServiceUnavailable, res.Code)
	}
}

func TestArticleListRequiresAuthenticatedOrApprovedLocalContext(t *testing.T) {
	handler := NewHandlerWithDependencies(Dependencies{
		ArticleReader: &fakeArticleReader{},
		UserContext: &fakeUserContextResolver{
			err: ErrUnauthenticated,
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusUnauthorized {
		t.Fatalf("expected %d, got %d", http.StatusUnauthorized, res.Code)
	}
}

func TestArticleListValidatesLimitBeforeReading(t *testing.T) {
	reader := &fakeArticleReader{}
	handler := NewHandlerWithDependencies(Dependencies{
		ArticleReader: reader,
		UserContext:   &fakeUserContextResolver{userID: "trusted-user"},
	})

	for _, path := range []string{
		"/api/v1/articles?limit=0",
		"/api/v1/articles?limit=101",
		"/api/v1/articles?limit=not-a-number",
		"/api/v1/articles?limit=1&limit=2",
	} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, req)

			if res.Code != http.StatusBadRequest {
				t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
			}
		})
	}
	if reader.callCount != 0 {
		t.Fatalf("reader must not be called for invalid limits; calls=%d", reader.callCount)
	}
}

func TestArticleListDependencyFailureDoesNotLeakInternalError(t *testing.T) {
	reader := &fakeArticleReader{err: errors.New("postgres secret details")}
	handler := NewHandlerWithDependencies(Dependencies{
		ArticleReader: reader,
		UserContext:   &fakeUserContextResolver{userID: "trusted-user"},
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected %d, got %d", http.StatusServiceUnavailable, res.Code)
	}
	if body := res.Body.String(); body == "" || strings.Contains(body, "postgres secret details") {
		t.Fatalf("internal error leaked in response: %q", body)
	}
}
