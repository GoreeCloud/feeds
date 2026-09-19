# GoreeCloud Feeds

GoreeCloud Feeds is the central project repository for the planned GoreeCloud feed aggregation, synchronization, reading, search, preservation, and content-management platform.

## Current state

**Repository state:** project and documentation foundation.

The five initial repositories exist, but the project does not yet have a verified deployable server, web client, stable protocol release, production deployment, or Stable release. Planned capabilities must not be read as implemented functionality.

## Repository role

This repository is the central coordination home for GoreeCloud Feeds. It owns project-level documentation and cross-repository coordination rather than the server or client implementation itself.

The initial repository set is:

- [feeds](https://github.com/GoreeCloud/feeds) — project overview, architecture, development, deployment, roadmap, compatibility, release, and cross-repository coordination.
- [feeds-server](https://github.com/GoreeCloud/feeds-server) — planned authoritative feed ingestion, processing, storage, search, synchronization, administration, and server API.
- [feeds-web](https://github.com/GoreeCloud/feeds-web) — planned Glaze UI web client.
- [feeds-protocol](https://github.com/GoreeCloud/feeds-protocol) — planned shared API, data, synchronization, event, error, and compatibility contracts.
- [feeds-shared](https://github.com/GoreeCloud/feeds-shared) — planned genuinely reusable internal models, validation, utilities, transformations, synchronization helpers, constants, and test fixtures.

Future desktop, mobile, documentation-site, and extension repositories remain deferred until their development phases are authorized.

## Architecture boundary

The planned architecture is client-server based. GoreeCloud Feeds Server is intended to remain authoritative for feed ingestion, processing, search, storage, synchronization, and administration. Clients are intended to communicate through a shared versioned protocol instead of independently redefining server state.

See [ARCHITECTURE.md](ARCHITECTURE.md).

## Project documentation

- [ARCHITECTURE.md](ARCHITECTURE.md) — current planned architecture and repository boundaries.
- [DEVELOPMENT.md](DEVELOPMENT.md) — development-state and repository workflow guidance.
- [DEPLOYMENT.md](DEPLOYMENT.md) — deployment boundary and current non-deployable status.
- [FEATURE-ROADMAP.md](FEATURE-ROADMAP.md) — repository-local roadmap index and implementation-state boundary.
- [COMPATIBILITY.md](COMPATIBILITY.md) — protocol and client/server compatibility principles.
- [COORDINATION.md](COORDINATION.md) — cross-repository ownership and synchronization rules.
- [RELEASES.md](RELEASES.md) — release-state boundaries and future release documentation.
- [CONTRIBUTING.md](CONTRIBUTING.md) — contribution and review workflow.

## Development and release status

No technology stack, package format, container image, public API version, supported client matrix, release candidate, production deployment, or Stable release is established by this repository foundation.

Repository and documentation changes should preserve that status distinction until implementation and authoritative verification support stronger claims.
