package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/GoreeCloud/feeds-server/internal/feed"
)

const (
	ProductName           = "GoreeCloud Feeds Server"
	APIVersion            = "v1"
	ProtocolVersion       = "0.1.0-dev"
	Lifecycle             = "development"
	ArticleListCapability = "articles:list-v1"
	defaultArticleLimit   = 50
	maxArticleLimit       = 100
)

var ErrUnauthenticated = errors.New("authenticated user context is unavailable")

type ArticleReader interface {
	ListRecentArticlesForUser(
		ctx context.Context,
		userID feed.ID,
		limit int,
	) ([]feed.UserArticleSummary, error)
}

type UserContextResolver interface {
	ResolveUserID(request *http.Request) (feed.ID, error)
}

type Dependencies struct {
	ArticleReader ArticleReader
	UserContext   UserContextResolver
}

type capabilityResponse struct {
	Product         string   `json:"product"`
	APIVersion      string   `json:"api_version"`
	ProtocolVersion string   `json:"protocol_version"`
	Lifecycle       string   `json:"lifecycle"`
	Capabilities    []string `json:"capabilities"`
}

type healthResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type articleSummaryResponse struct {
	ID          feed.ID    `json:"id"`
	FeedID      feed.ID    `json:"feed_id"`
	FeedTitle   string     `json:"feed_title"`
	URL         string     `json:"url"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	PublishedAt *time.Time `json:"published_at"`
	Summary     string     `json:"summary"`
	Language    string     `json:"language"`
	Read        bool       `json:"read"`
	Saved       bool       `json:"saved"`
	Favorite    bool       `json:"favorite"`
}

type articleListResponse struct {
	Articles []articleSummaryResponse `json:"articles"`
}

func NewHandler() http.Handler {
	return NewHandlerWithDependencies(Dependencies{})
}

func NewHandlerWithDependencies(deps Dependencies) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/capabilities", func(w http.ResponseWriter, _ *http.Request) {
		capabilities := make([]string, 0, 1)
		if deps.ArticleReader != nil && deps.UserContext != nil {
			capabilities = append(capabilities, ArticleListCapability)
		}
		writeJSON(w, http.StatusOK, capabilityResponse{
			Product:         ProductName,
			APIVersion:      APIVersion,
			ProtocolVersion: ProtocolVersion,
			Lifecycle:       Lifecycle,
			Capabilities:    capabilities,
		})
	})

	mux.HandleFunc("GET /api/v1/articles", func(w http.ResponseWriter, request *http.Request) {
		if deps.ArticleReader == nil || deps.UserContext == nil {
			writeJSON(w, http.StatusServiceUnavailable, errorResponse{
				Error:   "article_list_unavailable",
				Message: "Article listing is not configured for this server runtime.",
			})
			return
		}
		if request.URL.Query().Has("user_id") {
			writeJSON(w, http.StatusBadRequest, errorResponse{
				Error:   "user_context_must_be_server_derived",
				Message: "User context cannot be supplied by the article-list request.",
			})
			return
		}

		limit, err := articleLimit(request)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{
				Error:   "invalid_limit",
				Message: err.Error(),
			})
			return
		}

		userID, err := deps.UserContext.ResolveUserID(request)
		if err != nil {
			if errors.Is(err, ErrUnauthenticated) {
				writeJSON(w, http.StatusUnauthorized, errorResponse{
					Error:   "authentication_required",
					Message: "An authenticated or approved local user context is required.",
				})
				return
			}
			writeJSON(w, http.StatusServiceUnavailable, errorResponse{
				Error:   "user_context_unavailable",
				Message: "User context could not be resolved.",
			})
			return
		}
		if strings.TrimSpace(string(userID)) == "" {
			writeJSON(w, http.StatusServiceUnavailable, errorResponse{
				Error:   "user_context_unavailable",
				Message: "User context resolved to an invalid identifier.",
			})
			return
		}

		articles, err := deps.ArticleReader.ListRecentArticlesForUser(
			request.Context(),
			userID,
			limit,
		)
		if err != nil {
			writeJSON(w, http.StatusServiceUnavailable, errorResponse{
				Error:   "article_list_unavailable",
				Message: "Article listing is temporarily unavailable.",
			})
			return
		}

		response := articleListResponse{
			Articles: make([]articleSummaryResponse, 0, len(articles)),
		}
		for _, article := range articles {
			response.Articles = append(response.Articles, articleSummaryResponse{
				ID:          article.ArticleID,
				FeedID:      article.FeedID,
				FeedTitle:   article.FeedTitle,
				URL:         article.URL,
				Title:       article.Title,
				Author:      article.Author,
				PublishedAt: article.PublishedAt,
				Summary:     article.Summary,
				Language:    article.Language,
				Read:        article.Read,
				Saved:       article.Saved,
				Favorite:    article.Favorite,
			})
		}
		writeJSON(w, http.StatusOK, response)
	})

	mux.HandleFunc("GET /health/live", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
	})

	mux.HandleFunc("GET /health/ready", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
	})

	return mux
}

func articleLimit(request *http.Request) (int, error) {
	values := request.URL.Query()["limit"]
	if len(values) == 0 {
		return defaultArticleLimit, nil
	}
	if len(values) != 1 {
		return 0, errors.New("article list limit must be supplied at most once")
	}

	value, err := strconv.Atoi(values[0])
	if err != nil || value < 1 || value > maxArticleLimit {
		return 0, errors.New("article list limit must be an integer between 1 and 100")
	}
	return value, nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	_ = encoder.Encode(value)
}
