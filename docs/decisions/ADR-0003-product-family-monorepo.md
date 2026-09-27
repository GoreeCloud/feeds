# ADR-0003: Consolidate GoreeCloud Feeds into a product-family monorepo

## Status

Accepted — September 26, 2026

## Context

GoreeCloud Feeds began with five repositories: `feeds`, `feeds-server`, `feeds-web`, `feeds-protocol`, and `feeds-shared`.

Development quickly required coordinated server, protocol, web, and project-level changes. The active GoreeCloud Repository Architecture, Boundaries, and Organization standard identifies Feeds as a product-family monorepo candidate when those components evolve as one lifecycle.

The source repositories contain meaningful Development code, so current default-branch source must be preserved even though preserving every predecessor commit is not required for this migration.

## Decision

Use `GoreeCloud/feeds` as the canonical Feeds product-family monorepo with these boundaries:

- `services/server/` for the former `feeds-server` source;
- `apps/web/` for the former `feeds-web` source;
- `packages/protocol/` for the former `feeds-protocol` source;
- `packages/shared/` for the former `feeds-shared` source;
- root and `docs/` for project-family records;
- root `.github/workflows/` for monorepo-aware validation.

Current default-branch files from each component are migrated into those paths. Predecessor commit history is not imported into the monorepo commit graph.

## Alternatives Considered

### Keep the five repositories

Rejected because separation adds coordination and compatibility overhead without a sufficiently independent product lifecycle.

### Preserve all predecessor Git history through subtree/filter migration

Not selected for this migration because the owner explicitly accepted loss of non-meaningful historical data and prioritized beginning consolidation. Current live source is preserved.

## Consequences

- New Feeds-family development should occur in this repository.
- Server, web, and protocol changes can be reviewed atomically.
- Component CI remains separate but is executed from root workflows against monorepo paths.
- Predecessor repositories should be archived or clearly redirected after provider-level archival controls are available and the monorepo migration is verified.
- A component may be split out later only if it develops a genuinely independent lifecycle.
