# GoreeCloud Feeds — Changelogs

## 2026-10-04 — Bounded article-list source contract

- Added a bounded PostgreSQL chronological article read scoped to a server-supplied user and enabled subscriptions.
- Added the Development `GET /api/v1/articles` contract with `articles:list-v1` capability negotiation, 1..100 item bounds, and explicit 400/401/503 behavior.
- Added a dependency-gated Go HTTP handler that rejects client-supplied user identifiers and advertises article listing only when both reader and user-context resolver are configured.
- Added a strict TypeScript article-list client and deterministic server/web/protocol tests.
- The default runtime still does not wire PostgreSQL plus an accepted authenticated/local-only user-context resolver, so this milestone does not establish ordinary runtime article availability, deployment, Release Candidate, or Stable status.

## 2026-09-27 — Drive feature-roadmap migration

- Migrated feature-state authority to `IMPLEMENTED-FEATURES.md`, `PLANNED-FEATURES.md`, and this `CHANGELOGS.md`.
- Preserved the richer former Drive planning specification under `docs/history/drive-planned-features-source-2026-09-27.md` as historical, non-authoritative evidence.
- No lifecycle promotion, deployment, or Stable qualification is implied.
