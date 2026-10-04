package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/GoreeCloud/feeds-server/internal/feed"
)

func TestListRecentArticlesForUserIsBoundedChronologicalAndSubscriptionScoped(t *testing.T) {
	store := repositoryTestStore(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user := UserReference{ID: "list-user", IdentitySubject: "identity:list-user"}
	if err := store.UpsertUserReference(ctx, user); err != nil {
		t.Fatalf("create user: %v", err)
	}

	now := time.Date(2026, 10, 4, 17, 0, 0, 0, time.UTC)
	feeds := []struct {
		id       feed.ID
		title    string
		disabled bool
	}{
		{id: "list-feed-a", title: "Alpha Source"},
		{id: "list-feed-b", title: "Beta Source"},
		{id: "list-feed-disabled", title: "Disabled Source", disabled: true},
	}
	for _, item := range feeds {
		if err := store.UpsertFeed(ctx, FeedWrite{
			Feed: feed.Feed{
				ID:     item.id,
				URL:    "https://example.test/" + string(item.id) + ".xml",
				Title:  item.title,
				Source: feed.SourceInfo{Format: "rss"},
			},
			NormalizedURL: "https://example.test/" + string(item.id) + ".xml",
			SeenAt:        now,
		}); err != nil {
			t.Fatalf("create feed %s: %v", item.id, err)
		}
		if err := store.UpsertSubscription(ctx, SubscriptionWrite{
			Subscription: feed.Subscription{
				ID:        feed.ID("list-sub-" + string(item.id)),
				UserID:    user.ID,
				FeedID:    item.id,
				CreatedAt: now,
			},
			Disabled: item.disabled,
		}); err != nil {
			t.Fatalf("create subscription %s: %v", item.id, err)
		}
	}

	publishedOld := now.Add(-2 * time.Hour)
	publishedNew := now.Add(-time.Hour)

	articles := []ArticleWrite{
		{
			Article: feed.Article{
				ID:          "list-article-old",
				FeedID:      "list-feed-a",
				URL:         "https://example.test/articles/old",
				Title:       "Older article",
				Author:      "Example Author",
				PublishedAt: &publishedOld,
				Summary:     "Older summary",
				Language:    "en",
			},
			NormalizedURL:   "https://example.test/articles/old",
			RetrievedAt:     now.Add(-90 * time.Minute),
			SourceHistoryID: "list-history-old",
		},
		{
			Article: feed.Article{
				ID:          "list-article-new",
				FeedID:      "list-feed-b",
				URL:         "https://example.test/articles/new",
				Title:       "Newer article",
				PublishedAt: &publishedNew,
				Summary:     "Newer summary",
			},
			NormalizedURL:   "https://example.test/articles/new",
			RetrievedAt:     now.Add(-45 * time.Minute),
			SourceHistoryID: "list-history-new",
		},
		{
			Article: feed.Article{
				ID:      "list-article-undated",
				FeedID:  "list-feed-a",
				URL:     "https://example.test/articles/undated",
				Title:   "Recently retrieved undated article",
				Summary: "No publication time",
			},
			NormalizedURL:   "https://example.test/articles/undated",
			RetrievedAt:     now.Add(-30 * time.Minute),
			SourceHistoryID: "list-history-undated",
		},
		{
			Article: feed.Article{
				ID:          "list-article-disabled",
				FeedID:      "list-feed-disabled",
				URL:         "https://example.test/articles/disabled",
				Title:       "Must not be visible",
				PublishedAt: func() *time.Time { value := now; return &value }(),
			},
			NormalizedURL:   "https://example.test/articles/disabled",
			RetrievedAt:     now,
			SourceHistoryID: "list-history-disabled",
		},
	}
	for _, article := range articles {
		if err := store.UpsertArticle(ctx, article); err != nil {
			t.Fatalf("create article %s: %v", article.Article.ID, err)
		}
	}

	if err := store.UpsertArticleState(ctx, ArticleStateWrite{
		State: feed.ArticleState{
			UserID:    user.ID,
			ArticleID: "list-article-new",
			Read:      true,
			Saved:     true,
			UpdatedAt: now,
		},
	}); err != nil {
		t.Fatalf("store article state: %v", err)
	}

	recent, err := store.ListRecentArticlesForUser(ctx, user.ID, 2)
	if err != nil {
		t.Fatalf("ListRecentArticlesForUser() error = %v", err)
	}
	if len(recent) != 2 {
		t.Fatalf("expected 2 articles, got %d: %+v", len(recent), recent)
	}
	if recent[0].ArticleID != "list-article-undated" {
		t.Fatalf("unexpected first article: %+v", recent[0])
	}
	if recent[0].FeedTitle != "Alpha Source" || recent[0].Read || recent[0].Saved || recent[0].Favorite {
		t.Fatalf("unexpected default article state/source: %+v", recent[0])
	}
	if recent[1].ArticleID != "list-article-new" {
		t.Fatalf("unexpected second article: %+v", recent[1])
	}
	if recent[1].FeedTitle != "Beta Source" || !recent[1].Read || !recent[1].Saved || recent[1].Favorite {
		t.Fatalf("unexpected persisted article state/source: %+v", recent[1])
	}
	for _, article := range recent {
		if article.ArticleID == "list-article-disabled" {
			t.Fatal("article from disabled subscription must not be returned")
		}
	}
}

func TestListRecentArticlesForUserValidatesBounds(t *testing.T) {
	store := repositoryTestStore(t)
	ctx := context.Background()

	if _, err := store.ListRecentArticlesForUser(ctx, "", 10); err == nil {
		t.Fatal("expected blank user ID to fail")
	}
	if _, err := store.ListRecentArticlesForUser(ctx, "user", 0); err == nil {
		t.Fatal("expected zero limit to fail")
	}
	if _, err := store.ListRecentArticlesForUser(ctx, "user", 101); err == nil {
		t.Fatal("expected oversized limit to fail")
	}
}
