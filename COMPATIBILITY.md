# GoreeCloud Feeds Compatibility

## Current compatibility state

No Stable GoreeCloud Feeds protocol version, supported client matrix, or backward-compatibility guarantee has been established.

## Contract ownership

The shared Development protocol is maintained under `packages/protocol/` and defines the versioned client/server contract, data structures, errors, and compatibility rules.

Protocol-affecting changes should be made in `packages/protocol/` and validated in the same pull request as affected server or client updates when coordination is required.

Breaking changes must be explicit. Compatibility claims must be tied to verified implementation rather than inferred from directory names or roadmap intent.

## Current matrix

| Component | Verified compatibility status |
| --- | --- |
| `services/server/` | Development server source exists; no Stable protocol compatibility guarantee |
| `apps/web/` | Typed Development capabilities client exists; no Stable compatibility guarantee |
| `packages/protocol/` | Development OpenAPI contract exists; not Stable |
| Future clients | Not yet established for supported release use |
