# GoreeCloud Feeds Development

## Current development boundary

GoreeCloud Feeds is in controlled Development. The server, web client, protocol, and shared component boundaries now live in one product-family monorepo. No production deployment, Release Candidate, or Stable release is established.

Do not describe planned roadmap capabilities as implemented merely because they are documented.

## Component ownership

Use the path whose responsibility owns the change:

- repository root and `docs/` — product-family architecture, roadmap, compatibility, deployment, releases, contribution, and coordination.
- `services/server/` — server implementation.
- `apps/web/` — web client implementation.
- `packages/protocol/` — shared client/server contracts.
- `packages/shared/` — reusable implementation that is genuinely shared across components.

A coherent change may touch several component paths in one pull request when interfaces evolve together.

## Branch and pull-request workflow

Material changes should use short-lived topic branches created from the current intended base.

Preferred branch shape:

`<type>/<descriptive-kebab-case-subject>`

Use pull requests for material promotion to the authoritative default branch. Validate the exact candidate head and re-run affected validation if the candidate changes.

## Toolchain

Current Development selections include:

- server: Go 1.27.1;
- persistence target: PostgreSQL 18;
- web development/CI runtime: Node.js 24.21.0 LTS;
- web language: TypeScript 7.0.2;
- shared client/server contract: HTTP/JSON described with OpenAPI 3.1;
- API path versioning beginning under `/api/v1/`, with v1 still Development rather than Stable.

## Validation

Root GitHub workflows validate:

- the Go server from `services/server/`;
- the TypeScript web client from `apps/web/`;
- the protocol contract from `packages/protocol/`.

Passing validation is Development evidence only.

## Secrets and local state

Do not commit reusable credentials, private keys, active tokens, production configuration, personal data, or protected operational data.
