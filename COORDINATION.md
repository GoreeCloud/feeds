# GoreeCloud Feeds Cross-Repository Coordination

## Purpose

This document defines the current project-level ownership map for the five initial GoreeCloud Feeds repositories.

## Ownership map

| Repository | Primary responsibility |
| --- | --- |
| `feeds` | Project-wide documentation, architecture, roadmap, deployment, compatibility, release, contribution, and coordination |
| `feeds-server` | Authoritative service and server-side product behavior |
| `feeds-web` | Glaze UI web client |
| `feeds-protocol` | Shared client/server contracts and compatibility rules |
| `feeds-shared` | Reusable internal implementation shared by more than one repository |

## Dependency direction

Protocol contracts should be owned by `feeds-protocol`.

Server-specific code should remain in `feeds-server`.

Web-specific code should remain in `feeds-web`.

Reusable code should move to `feeds-shared` only when actual reuse justifies the shared ownership boundary.

Project-wide documentation belongs in `feeds` unless a document describes repository-local implementation details.

## Coordinated changes

For a cross-repository change:

1. identify the authoritative owner of the contract or behavior;
2. establish the owning change first;
3. record compatibility or migration implications;
4. update dependent repositories against the verified owner state;
5. validate each repository independently; and
6. update project-level documentation only after the relevant authoritative state is known.

## Current state

The five repositories currently represent repository boundaries, not a verified integrated product. No cross-repository compatibility or runtime integration is established by repository existence alone.
