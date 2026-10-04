import assert from "node:assert/strict";
import {
  ARTICLE_LIST_CAPABILITY,
  DEFAULT_ARTICLE_LIST_LIMIT,
  fetchArticles,
  parseArticleListResponse,
} from "../dist/protocol/articles.js";
import { ProtocolContractError } from "../dist/protocol/capabilities.js";

const validArticle = Object.freeze({
  id: "article-1",
  feed_id: "feed-1",
  feed_title: "Example Feed",
  url: "https://example.test/articles/1",
  title: "Article One",
  author: "Example Author",
  published_at: "2026-10-04T16:00:00Z",
  summary: "Summary",
  language: "en",
  read: false,
  saved: true,
  favorite: false,
});

assert.equal(ARTICLE_LIST_CAPABILITY, "articles:list-v1");
assert.equal(DEFAULT_ARTICLE_LIST_LIMIT, 50);
assert.deepEqual(parseArticleListResponse({ articles: [validArticle] }), {
  articles: [validArticle],
});

assert.throws(
  () => parseArticleListResponse({ articles: [{ ...validArticle, user_id: "not-allowed" }] }),
  ProtocolContractError,
);
assert.throws(
  () => parseArticleListResponse({ articles: [{ ...validArticle, published_at: "not-a-date" }] }),
  ProtocolContractError,
);
assert.throws(
  () => parseArticleListResponse({ articles: [{ ...validArticle, read: "false" }] }),
  ProtocolContractError,
);
assert.throws(
  () => parseArticleListResponse({ articles: Array.from({ length: 101 }, () => validArticle) }),
  ProtocolContractError,
);

let requestedUrl;
let requestedInit;
const response = await fetchArticles(
  "http://127.0.0.1:8080/some/base",
  25,
  async (input, init) => {
    requestedUrl = String(input);
    requestedInit = init;
    return new Response(JSON.stringify({ articles: [validArticle] }), {
      status: 200,
      headers: { "content-type": "application/json; charset=utf-8" },
    });
  },
);

assert.deepEqual(response, { articles: [validArticle] });
assert.equal(requestedUrl, "http://127.0.0.1:8080/api/v1/articles?limit=25");
assert.equal(requestedInit.method, "GET");
assert.equal(requestedInit.credentials, "include");
assert.equal(requestedInit.redirect, "error");
assert.equal(requestedInit.referrerPolicy, "no-referrer");
assert.equal(requestedInit.cache, "no-store");
assert.equal(requestedInit.headers.Accept, "application/json");

await assert.rejects(
  fetchArticles("http://127.0.0.1:8080/", 0, async () => {
    throw new Error("fetcher should not be called");
  }),
  ProtocolContractError,
);
await assert.rejects(
  fetchArticles("http://127.0.0.1:8080/", 101, async () => {
    throw new Error("fetcher should not be called");
  }),
  ProtocolContractError,
);
await assert.rejects(
  fetchArticles("ftp://127.0.0.1/", 10, async () => {
    throw new Error("fetcher should not be called");
  }),
  ProtocolContractError,
);
await assert.rejects(
  fetchArticles("http://user:password@127.0.0.1:8080/", 10, async () => {
    throw new Error("fetcher should not be called");
  }),
  ProtocolContractError,
);
await assert.rejects(
  fetchArticles("http://127.0.0.1:8080/", 10, async () =>
    new Response(JSON.stringify({ error: "authentication_required" }), {
      status: 401,
      headers: { "content-type": "application/json" },
    }),
  ),
  ProtocolContractError,
);

console.log("GoreeCloud Feeds article protocol client validation passed.");
