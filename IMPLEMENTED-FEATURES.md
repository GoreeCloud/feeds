# GoreeCloud Feeds — Implemented Features

> **Authority:** Repository-native implemented-feature record.  
> **Boundary:** This records only verified Development source foundations already documented by the repository. It does not establish feature completion, production deployment, Release Candidate, or Stable status.

## Current Development source foundations

GoreeCloud Feeds remains in **Development** and has no Stable release or production deployment.

Verified source foundations now present in this repository include:

- Go server runtime and tests.
- Normalized feed/article models plus bounded RSS/Atom parsing.
- Conservative source-scoped article deduplication.
- PostgreSQL 18 schema/migration, connectivity, and durable core repository foundations.
- Bounded remote-feed retrieval with destination/redirect controls, conditional requests, concurrency limits, transient retry/backoff, and bounded `Retry-After` handling.
- Development OpenAPI contract under `packages/protocol/`.
- TypeScript web toolchain and typed Development capabilities client under `apps/web/`.
- Monorepo CI that validates the server, web client, and protocol from their new paths.

The network-visible server runtime is not yet fully wired to the persistence and retrieval pipelines, and production authentication, synchronization, full Glaze UI product experience, deployment packaging, backup/recovery acceptance, Release Candidate, and Stable qualification remain incomplete.
