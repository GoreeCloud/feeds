export const FEEDS_PROTOCOL_VERSION = "0.1.0-dev" as const;
export const FEEDS_API_VERSION = "v1" as const;
export const FEEDS_SERVER_PRODUCT = "GoreeCloud Feeds Server" as const;
export const FEEDS_SERVER_LIFECYCLE = "development" as const;

export interface CapabilityResponse {
  product: typeof FEEDS_SERVER_PRODUCT;
  api_version: typeof FEEDS_API_VERSION;
  protocol_version: typeof FEEDS_PROTOCOL_VERSION;
  lifecycle: typeof FEEDS_SERVER_LIFECYCLE;
  capabilities: string[];
}

export class ProtocolContractError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ProtocolContractError";
  }
}

type FetchLike = (
  input: RequestInfo | URL,
  init?: RequestInit,
) => Promise<Response>;

const expectedKeys = new Set([
  "product",
  "api_version",
  "protocol_version",
  "lifecycle",
  "capabilities",
]);

function isRecord(value: unknown): value is Record<string, unknown> {
  return typeof value === "object" && value !== null && !Array.isArray(value);
}

function requireExactKeys(value: Record<string, unknown>): void {
  const keys = Object.keys(value);
  if (keys.length !== expectedKeys.size || keys.some((key) => !expectedKeys.has(key))) {
    throw new ProtocolContractError("Capability response contains unexpected or missing fields.");
  }
}

export function parseCapabilityResponse(value: unknown): CapabilityResponse {
  if (!isRecord(value)) {
    throw new ProtocolContractError("Capability response must be a JSON object.");
  }

  requireExactKeys(value);

  if (value.product !== FEEDS_SERVER_PRODUCT) {
    throw new ProtocolContractError("Capability response product does not match GoreeCloud Feeds Server.");
  }
  if (value.api_version !== FEEDS_API_VERSION) {
    throw new ProtocolContractError("Capability response API version is incompatible.");
  }
  if (value.protocol_version !== FEEDS_PROTOCOL_VERSION) {
    throw new ProtocolContractError("Capability response protocol version is incompatible.");
  }
  if (value.lifecycle !== FEEDS_SERVER_LIFECYCLE) {
    throw new ProtocolContractError("Capability response lifecycle is incompatible.");
  }
  if (
    !Array.isArray(value.capabilities) ||
    value.capabilities.some((capability) => typeof capability !== "string")
  ) {
    throw new ProtocolContractError("Capability response capabilities must be an array of strings.");
  }

  return {
    product: FEEDS_SERVER_PRODUCT,
    api_version: FEEDS_API_VERSION,
    protocol_version: FEEDS_PROTOCOL_VERSION,
    lifecycle: FEEDS_SERVER_LIFECYCLE,
    capabilities: [...value.capabilities],
  };
}

function capabilityEndpoint(baseUrl: string | URL): URL {
  const base = new URL(baseUrl);
  if (base.protocol !== "http:" && base.protocol !== "https:") {
    throw new ProtocolContractError("Feeds Server base URL must use HTTP or HTTPS.");
  }
  if (base.username || base.password) {
    throw new ProtocolContractError("Feeds Server base URL must not embed credentials.");
  }

  return new URL("/api/v1/capabilities", base);
}

export async function fetchCapabilities(
  baseUrl: string | URL,
  fetcher: FetchLike = fetch,
): Promise<CapabilityResponse> {
  const endpoint = capabilityEndpoint(baseUrl);
  const response = await fetcher(endpoint, {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
    cache: "no-store",
    credentials: "omit",
    redirect: "error",
    referrerPolicy: "no-referrer",
  });

  if (!response.ok) {
    throw new ProtocolContractError(
      `Capability request failed with HTTP ${response.status}.`,
    );
  }

  const contentType = response.headers.get("content-type")?.toLowerCase() ?? "";
  if (!contentType.startsWith("application/json")) {
    throw new ProtocolContractError("Capability response must use application/json.");
  }

  let payload: unknown;
  try {
    payload = await response.json();
  } catch {
    throw new ProtocolContractError("Capability response contains invalid JSON.");
  }

  return parseCapabilityResponse(payload);
}
