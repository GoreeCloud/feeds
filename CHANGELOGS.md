# GoreeCloud Feeds — Changelogs

## 2026-10-04 — Optional Development PostgreSQL runtime wiring

- Added optional `GOREECLOUD_FEEDS_DATABASE_URL` startup wiring for the Development server.
- When configured, the server opens the bounded PostgreSQL pool, applies embedded migrations before serving, and injects the store as the article reader.
- When unconfigured, the existing dependency-free capability/health Development runtime remains available.
- No user-context resolver is configured by this tranche; `articles:list-v1` therefore remains unadvertised and `GET /api/v1/articles` remains fail-closed.
- CI verifies configured PostgreSQL 18.6 startup/migration behavior and verifies that connection errors do not echo the configured password.
- This does not establish authenticated/local-only identity, live article availability, deployment, Release Candidate, or Stable status.

## 2026-10-04 — Bounded article-list source contract

- Added a bounded PostgreSQL chronological article read scoped to a server-supplied user and enabled subscriptions.
- Added the Development `GET /api/v1/articles` contract with `articles:list-v1` capability negotiation, 1..100 item bounds, and explicit 400/401/503 behavior.
- Added a dependency-gated Go HTTP handler that rejects client-supplied user identifiers and advertises article listing only when both reader and user-context resolver are configured.
- Added a strict TypeScript article-list client and deterministic server/web/protocol tests.
- The article contract milestone did not by itself establish runtime database or identity wiring; later Development work added optional PostgreSQL startup while authenticated/local-only user-context resolution remains open.

## 2026-09-27 — Drive feature-roadmap migration

- Migrated feature-state authority to `IMPLEMENTED-FEATURES.md`, `PLANNED-FEATURES.md`, and this `CHANGELOGS.md`.
- Preserved the richer former Drive planning specification under `docs/history/drive-planned-features-source-2026-09-27.md` as historical, non-authoritative evidence.
- No lifecycle promotion, deployment, or Stable qualification is implied.
