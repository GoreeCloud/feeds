# GoreeCloud Feeds Component Coordination

## Purpose

This document defines ownership boundaries inside the GoreeCloud Feeds product-family monorepo.

## Ownership map

| Path | Primary responsibility |
| --- | --- |
| repository root | Product-family architecture, roadmap, deployment, compatibility, releases, contribution, and coordination |
| `services/server/` | Authoritative service and server-side product behavior |
| `apps/web/` | Glaze UI web-client implementation |
| `packages/protocol/` | Shared client/server contracts and compatibility rules |
| `packages/shared/` | Reusable implementation shared by more than one Feeds component |

## Dependency direction

Protocol-affecting changes belong in `packages/protocol/`.

Server-specific code belongs in `services/server/`.

Web-specific code belongs in `apps/web/`.

Reusable implementation should move into `packages/shared/` only when actual reuse justifies a shared ownership boundary.

Project-wide documentation belongs at the repository root or under `docs/`; component-local implementation records stay with their component.

## Coordinated changes

A change may update multiple component paths in one pull request when the behavior or contract evolves together. The change must still identify the authoritative owner of each contract, record compatibility implications, and run the affected component validation.

The monorepo removes cross-repository synchronization overhead; it does not erase component, runtime, API, release-artifact, or security boundaries.

## Current state

The former `feeds-server`, `feeds-web`, `feeds-protocol`, and `feeds-shared` source boundaries are represented by internal monorepo paths. Their predecessor repositories are migration predecessors rather than the intended location for new Feeds-family development once this migration is merged.
