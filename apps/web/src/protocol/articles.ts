import { ProtocolContractError } from "./capabilities.js";

export const ARTICLE_LIST_CAPABILITY = "articles:list-v1" as const;
export const DEFAULT_ARTICLE_LIST_LIMIT = 50;
export const MAX_ARTICLE_LIST_LIMIT = 100;

export interface ArticleSummary {
  id: string;
  feed_id: string;
  feed_title: string;
  url: string;
  title: string;
  author: string;
  published_at: string | null;
  summary: string;
  language: string;
  read: boolean;
  saved: boolean;
  favorite: boolean;
}

export interface ArticleListResponse {
  articles: ArticleSummary[];
}

type FetchLike = (
  input: RequestInfo | URL,
  init?: RequestInit,
) => Promise<Response>;

const articleKeys = new Set([
  "id",
  "feed_id",
  "feed_title",
  "url",
  "title",
  "author",
  "published_at",
  "summary",
  "language",
  "read",
  "saved",
  "favorite",
]);

function isRecord(value: unknown): value is Record<string, unknown> {
  return typeof value === "object" && value !== null && !Array.isArray(value);
}

function requireExactKeys(
  value: Record<string, unknown>,
  expected: ReadonlySet<string>,
  subject: string,
): void {
  const keys = Object.keys(value);
  if (keys.length !== expected.size || keys.some((key) => !expected.has(key))) {
    throw new ProtocolContractError(
      `${subject} contains unexpected or missing fields.`,
    );
  }
}

function requireString(
  value: Record<string, unknown>,
  key: string,
  subject: string,
): string {
  const field = value[key];
  if (typeof field !== "string") {
    throw new ProtocolContractError(`${subject} ${key} must be a string.`);
  }
  return field;
}

function requireBoolean(
  value: Record<string, unknown>,
  key: string,
  subject: string,
): boolean {
  const field = value[key];
  if (typeof field !== "boolean") {
    throw new ProtocolContractError(`${subject} ${key} must be a boolean.`);
  }
  return field;
}

function parseArticleSummary(value: unknown): ArticleSummary {
  if (!isRecord(value)) {
    throw new ProtocolContractError("Article summary must be a JSON object.");
  }
  requireExactKeys(value, articleKeys, "Article summary");

  const id = requireString(value, "id", "Article summary");
  const feedId = requireString(value, "feed_id", "Article summary");
  if (!id.trim() || !feedId.trim()) {
    throw new ProtocolContractError("Article summary id and feed_id must not be blank.");
  }

  const publishedAt = value.published_at;
  if (publishedAt !== null && typeof publishedAt !== "string") {
    throw new ProtocolContractError(
      "Article summary published_at must be a string or null.",
    );
  }
  if (
    typeof publishedAt === "string" &&
    (!publishedAt.trim() || Number.isNaN(Date.parse(publishedAt)))
  ) {
    throw new ProtocolContractError(
      "Article summary published_at must be a valid date-time when present.",
    );
  }

  return {
    id,
    feed_id: feedId,
    feed_title: requireString(value, "feed_title", "Article summary"),
    url: requireString(value, "url", "Article summary"),
    title: requireString(value, "title", "Article summary"),
    author: requireString(value, "author", "Article summary"),
    published_at: publishedAt,
    summary: requireString(value, "summary", "Article summary"),
    language: requireString(value, "language", "Article summary"),
    read: requireBoolean(value, "read", "Article summary"),
    saved: requireBoolean(value, "saved", "Article summary"),
    favorite: requireBoolean(value, "favorite", "Article summary"),
  };
}

export function parseArticleListResponse(value: unknown): ArticleListResponse {
  if (!isRecord(value)) {
    throw new ProtocolContractError("Article list response must be a JSON object.");
  }
  requireExactKeys(value, new Set(["articles"]), "Article list response");

  if (!Array.isArray(value.articles)) {
    throw new ProtocolContractError("Article list articles must be an array.");
  }
  if (value.articles.length > MAX_ARTICLE_LIST_LIMIT) {
    throw new ProtocolContractError("Article list exceeds the maximum supported size.");
  }

  return {
    articles: value.articles.map(parseArticleSummary),
  };
}

function articleEndpoint(baseUrl: string | URL, limit: number): URL {
  if (!Number.isInteger(limit) || limit < 1 || limit > MAX_ARTICLE_LIST_LIMIT) {
    throw new ProtocolContractError("Article list limit must be an integer between 1 and 100.");
  }

  const base = new URL(baseUrl);
  if (base.protocol !== "http:" && base.protocol !== "https:") {
    throw new ProtocolContractError("Feeds Server base URL must use HTTP or HTTPS.");
  }
  if (base.username || base.password) {
    throw new ProtocolContractError("Feeds Server base URL must not embed credentials.");
  }

  const endpoint = new URL("/api/v1/articles", base);
  endpoint.searchParams.set("limit", String(limit));
  return endpoint;
}

export async function fetchArticles(
  baseUrl: string | URL,
  limit = DEFAULT_ARTICLE_LIST_LIMIT,
  fetcher: FetchLike = fetch,
): Promise<ArticleListResponse> {
  const endpoint = articleEndpoint(baseUrl, limit);
  const response = await fetcher(endpoint, {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
    cache: "no-store",
    credentials: "include",
    redirect: "error",
    referrerPolicy: "no-referrer",
  });

  if (!response.ok) {
    throw new ProtocolContractError(
      `Article list request failed with HTTP ${response.status}.`,
    );
  }

  const contentType = response.headers.get("content-type")?.toLowerCase() ?? "";
  if (!contentType.startsWith("application/json")) {
    throw new ProtocolContractError("Article list response must use application/json.");
  }

  let payload: unknown;
  try {
    payload = await response.json();
  } catch {
    throw new ProtocolContractError("Article list response contains invalid JSON.");
  }

  return parseArticleListResponse(payload);
}
