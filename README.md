# GoreeCloud Feeds

GoreeCloud Feeds is the GoreeCloud feed aggregation, synchronization, reading, search, preservation, and content-management platform.

This repository is the canonical **Feeds product-family monorepo**. The server, web client, protocol contract, reusable shared package boundary, and project-wide documentation are maintained together while retaining clear internal ownership boundaries.

## Repository layout

- `services/server/` — Go server runtime, feed processing, retrieval, and PostgreSQL persistence foundations.
- `apps/web/` — TypeScript web-client foundation and typed Development protocol client.
- `packages/protocol/` — versioned Development API contract and validation tooling.
- `packages/shared/` — genuinely reusable Feeds code and shared technical records when reuse is justified.
- `docs/` — product-family architecture decisions and supporting documentation.
- `.github/workflows/` — monorepo-aware server, web, and protocol validation.

## Current Development state

GoreeCloud Feeds remains in **Development** and has no Stable release or production deployment.

Verified source foundations now present in this repository include:

- Go server runtime and tests.
- Normalized feed/article models plus bounded RSS/Atom parsing.
- Conservative source-scoped article deduplication.
- PostgreSQL 18 schema/migration, connectivity, durable core repository foundations, bounded chronological article reads scoped to a server-supplied user, and optional Development runtime startup/migration wiring through `GOREECLOUD_FEEDS_DATABASE_URL`.
- Bounded remote-feed retrieval with destination/redirect controls, conditional requests, concurrency limits, transient retry/backoff, and bounded `Retry-After` handling.
- Development OpenAPI contract under `packages/protocol/`, including capability discovery and bounded `GET /api/v1/articles` with server-derived user context.
- TypeScript web toolchain plus strict Development capability and article-list clients under `apps/web/`.
- Monorepo CI that validates the server, web client, and protocol from their new paths.

The Development runtime can now open and migrate PostgreSQL when `GOREECLOUD_FEEDS_DATABASE_URL` is configured, injecting that store as the article reader. The article-list handler remains dependency-gated and fails closed because no accepted authenticated/local-only user-context resolver is configured; `articles:list-v1` therefore remains unadvertised in the default runtime. Retrieval/runtime orchestration, production authentication, synchronization, full Glaze UI product experience, deployment packaging, backup/recovery acceptance, Release Candidate, and Stable qualification remain incomplete.

## Architecture

The server remains authoritative for server-side feed and account state. Clients consume versioned protocol contracts and maintain only the local state required for interface behavior, caching, offline operation, and synchronization.

See `ARCHITECTURE.md` and the decisions under `docs/decisions/`.

## Development and release integrity

A source file, passing workflow, merged pull request, or monorepo migration does not itself establish a release or production acceptance. Current implementation claims must remain tied to verified source and test evidence.

## License

Component license files migrated from the predecessor repositories remain alongside their component sources. Project-wide licensing must continue to follow the applicable GoreeCloud licensing governance.
