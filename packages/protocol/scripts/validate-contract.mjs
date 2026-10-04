import { readFile } from "node:fs/promises";

const path = new URL("../openapi/feeds-v1.json", import.meta.url);
const raw = await readFile(path, "utf8");
const doc = JSON.parse(raw);

const fail = (message) => {
  console.error(message);
  process.exitCode = 1;
};

if (doc.openapi !== "3.1.0") fail("Expected OpenAPI 3.1.0.");
if (doc.info?.version !== "0.1.0-dev") fail("Expected Development protocol version 0.1.0-dev.");

const capabilityPath = doc.paths?.["/api/v1/capabilities"]?.get;
if (!capabilityPath) fail("Missing GET /api/v1/capabilities.");

const schema = doc.components?.schemas?.CapabilityResponse;
if (!schema) fail("Missing CapabilityResponse schema.");

const required = new Set(schema?.required ?? []);
for (const field of ["product", "api_version", "protocol_version", "lifecycle", "capabilities"]) {
  if (!required.has(field)) fail(`CapabilityResponse must require ${field}.`);
}

if (schema?.properties?.api_version?.const !== "v1") fail("API version must remain v1.");
if (schema?.properties?.protocol_version?.const !== "0.1.0-dev") fail("Protocol version constant mismatch.");
if (!schema?.properties?.lifecycle?.enum?.includes("development")) fail("Lifecycle must include development.");

const articlePath = doc.paths?.["/api/v1/articles"]?.get;
if (!articlePath) fail("Missing GET /api/v1/articles.");

const articleLimit = articlePath?.parameters?.find?.(
  (parameter) => parameter?.name === "limit" && parameter?.in === "query",
);
if (!articleLimit) fail("Article list must define the limit query parameter.");
if (articleLimit?.schema?.minimum !== 1 || articleLimit?.schema?.maximum !== 100) {
  fail("Article list limit must remain bounded to 1..100.");
}
if (articlePath?.parameters?.some?.((parameter) => parameter?.name === "user_id")) {
  fail("Article list must not accept a client-supplied user_id.");
}

for (const status of ["200", "400", "401", "503"]) {
  if (!articlePath?.responses?.[status]) {
    fail(`Article list is missing required ${status} response.`);
  }
}

const articleSchema = doc.components?.schemas?.ArticleSummary;
if (!articleSchema) fail("Missing ArticleSummary schema.");
const articleRequired = new Set(articleSchema?.required ?? []);
for (const field of [
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
]) {
  if (!articleRequired.has(field)) fail(`ArticleSummary must require ${field}.`);
}

const articleListSchema = doc.components?.schemas?.ArticleListResponse;
if (!articleListSchema) fail("Missing ArticleListResponse schema.");
if (articleListSchema?.properties?.articles?.maxItems !== 100) {
  fail("ArticleListResponse must remain bounded to 100 articles.");
}

const errorSchema = doc.components?.schemas?.ErrorResponse;
if (!errorSchema) fail("Missing ErrorResponse schema.");

if (!process.exitCode) {
  console.log("GoreeCloud Feeds Development contract validation passed.");
}
