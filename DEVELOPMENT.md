# GoreeCloud Feeds Development

## Current development boundary

The project has entered its first controlled Development implementation tranche. The five initial repositories have governed foundations and ADR-0001 selects the initial toolchain, but no deployable application, production service, Release Candidate, or Stable release is established.

Do not describe planned roadmap capabilities as implemented merely because they are documented.

## Repository ownership

Use the repository whose responsibility owns the change:

- `feeds` — project-wide architecture, roadmap, compatibility, deployment, release, contribution, and coordination documentation.
- `feeds-server` — server implementation.
- `feeds-web` — web client implementation.
- `feeds-protocol` — shared client/server contracts.
- `feeds-shared` — reusable implementation that is genuinely shared across repositories.

Avoid duplicating implementation across repositories when a clear owner exists.

## Branch and pull-request workflow

Material changes should use short-lived topic branches created from the current intended base.

Preferred branch shape:

`<type>/<descriptive-kebab-case-subject>`

Examples include `feature/`, `fix/`, `docs/`, `test/`, `ci/`, `build/`, `chore/`, and `security/`.

Use pull requests for material promotion to the authoritative default branch. Validate the exact candidate head, and re-run affected validation if the candidate changes after review.

## Status integrity

A branch, pull request, successful tool call, merge, package, or repository file does not by itself prove release, deployment, production acceptance, or Stable status.

Documentation must distinguish:

- planned;
- implemented;
- verified;
- released;
- deployed; and
- production-accepted

states when those distinctions are material.

## Toolchain selection

ADR-0001 selects the initial Development toolchain:

- server: Go 1.27.1;
- web development/CI runtime: Node.js 24.21.0 LTS;
- web language: TypeScript 7.0.2;
- shared client/server contract: REST-style HTTP/JSON described with OpenAPI 3.1;
- API path versioning begins under `/api/v1/`, with v1 still Development rather than Stable;
- feeds-shared intentionally has no selected implementation language/package until genuine multi-repository reuse exists.

Database technology, authentication implementation, web UI framework, container/deployment model, and production packaging remain deferred decisions.

## Secrets and local state

Do not commit reusable credentials, private keys, active tokens, production configuration, personal data, or other protected values.

Repository-specific ignore and environment rules should be added when each implementation toolchain is selected.
