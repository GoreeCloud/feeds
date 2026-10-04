# GoreeCloud Feeds — Implemented Features

> **Authority:** Repository-native implemented-feature record.  
> **Boundary:** This records only verified Development source foundations already documented by the repository. It does not establish feature completion, production deployment, Release Candidate, or Stable status.

## Current Development source foundations

GoreeCloud Feeds remains in **Development** and has no Stable release or production deployment.

Verified source foundations now present in this repository include:

- Go server runtime and tests.
- Normalized feed/article models plus bounded RSS/Atom parsing.
- Conservative source-scoped article deduplication.
- PostgreSQL 18 schema/migration, connectivity, durable core repository foundations, bounded chronological article reads scoped to a server-supplied user, and optional Development runtime startup/migration wiring through `GOREECLOUD_FEEDS_DATABASE_URL`.
- Bounded remote-feed retrieval with destination/redirect controls, conditional requests, concurrency limits, transient retry/backoff, and bounded `Retry-After` handling.
- Development OpenAPI contract under `packages/protocol/`, including `GET /api/v1/capabilities` and bounded `GET /api/v1/articles`.
- TypeScript web toolchain plus strict Development capability and article-list clients under `apps/web/`.
- Monorepo CI that validates the server, web client, and protocol from their new paths.

The article-list endpoint is source-implemented behind dependency injection and capability gating. The Development runtime can open and migrate PostgreSQL and inject the store as the article reader, but it still lacks an accepted authenticated/local-only user-context resolver, so `articles:list-v1` remains unadvertised and ordinary runtime article availability is not established. Production authentication, synchronization, full Glaze UI product experience, deployment packaging, backup/recovery acceptance, Release Candidate, and Stable qualification remain incomplete.
