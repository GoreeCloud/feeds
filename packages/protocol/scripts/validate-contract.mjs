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

if (!process.exitCode) {
  console.log("GoreeCloud Feeds Development contract validation passed.");
}
