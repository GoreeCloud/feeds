import assert from "node:assert/strict";
import {
  FEEDS_API_VERSION,
  FEEDS_PROTOCOL_VERSION,
  FEEDS_SERVER_LIFECYCLE,
  FEEDS_SERVER_PRODUCT,
  ProtocolContractError,
  fetchCapabilities,
  parseCapabilityResponse,
} from "../dist/protocol/capabilities.js";

const valid = Object.freeze({
  product: FEEDS_SERVER_PRODUCT,
  api_version: FEEDS_API_VERSION,
  protocol_version: FEEDS_PROTOCOL_VERSION,
  lifecycle: FEEDS_SERVER_LIFECYCLE,
  capabilities: [],
});

assert.deepEqual(parseCapabilityResponse(valid), valid);

assert.throws(
  () => parseCapabilityResponse({ ...valid, unexpected: true }),
  ProtocolContractError,
);
assert.throws(
  () => parseCapabilityResponse({ ...valid, protocol_version: "0.2.0-dev" }),
  ProtocolContractError,
);
assert.throws(
  () => parseCapabilityResponse({ ...valid, capabilities: [42] }),
  ProtocolContractError,
);

let requestedUrl;
let requestedInit;
const result = await fetchCapabilities(
  "http://127.0.0.1:8080/some/base",
  async (input, init) => {
    requestedUrl = String(input);
    requestedInit = init;
    return new Response(JSON.stringify(valid), {
      status: 200,
      headers: { "content-type": "application/json; charset=utf-8" },
    });
  },
);

assert.deepEqual(result, valid);
assert.equal(requestedUrl, "http://127.0.0.1:8080/api/v1/capabilities");
assert.equal(requestedInit.method, "GET");
assert.equal(requestedInit.credentials, "omit");
assert.equal(requestedInit.redirect, "error");
assert.equal(requestedInit.referrerPolicy, "no-referrer");
assert.equal(requestedInit.cache, "no-store");
assert.equal(requestedInit.headers.Accept, "application/json");

await assert.rejects(
  fetchCapabilities("ftp://127.0.0.1/", async () => {
    throw new Error("fetcher should not be called");
  }),
  ProtocolContractError,
);

await assert.rejects(
  fetchCapabilities("http://user:password@127.0.0.1:8080/", async () => {
    throw new Error("fetcher should not be called");
  }),
  ProtocolContractError,
);

await assert.rejects(
  fetchCapabilities("http://127.0.0.1:8080/", async () =>
    new Response("not json", {
      status: 200,
      headers: { "content-type": "text/plain" },
    }),
  ),
  ProtocolContractError,
);

await assert.rejects(
  fetchCapabilities("http://127.0.0.1:8080/", async () =>
    new Response(JSON.stringify(valid), {
      status: 503,
      headers: { "content-type": "application/json" },
    }),
  ),
  ProtocolContractError,
);

console.log("GoreeCloud Feeds protocol client validation passed.");
