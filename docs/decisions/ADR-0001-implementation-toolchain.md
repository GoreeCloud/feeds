# ADR-0001 — Initial GoreeCloud Feeds Implementation Toolchain

## Status

Accepted for the initial Development foundation.

This decision establishes the first implementation toolchain. It does not establish a Stable API, production deployment, permanent database choice, authentication implementation, or final packaging model.

## Decision date

2026-09-19

## Context

GoreeCloud Feeds requires:

- a long-running self-hosted authoritative server;
- a browser-based GoreeCloud web client;
- explicit, versioned client/server contracts;
- low dependency burden;
- reproducible builds;
- strong portability and long-term maintainability;
- current Glaze UI adoption on web surfaces; and
- clear separation between protocol authority, server implementation, web implementation, and genuinely shared code.

GoreeCloud programming-language strategy identifies Go as the primary infrastructure/network-service language and TypeScript as the primary web-application/interface language. GoreeCloud API strategy prefers documented HTTP APIs using common web standards, JSON where appropriate, explicit versioning, and OpenAPI when practical.

## Decision

### Server

Use Go for GoreeCloud Feeds Server.

Initial toolchain:

- Go 1.27.1;
- standard-library net/http for the first bounded server foundation;
- standard-library testing for the first validation slice;
- no external runtime dependency in the first server tranche.

The initial server will expose only development-safe capability discovery plus liveness/readiness endpoints. Feed ingestion, storage, search, synchronization, authentication, and administration remain separate implementation tranches.

### Web

Use TypeScript for GoreeCloud Feeds Web.

Initial toolchain:

- Node.js 24.21.0 LTS for reproducible development/CI execution;
- TypeScript 7.0.2;
- no UI framework selected in ADR-0001;
- no production browser application rendered in the first toolchain tranche.

A UI framework may be selected later if it materially improves accessibility, maintainability, testing, performance, and Glaze UI integration without unnecessary lock-in. Any rendered Feeds Web interface must use the current approved production Glaze UI; current authority is Glaze UI V1.6.0.

### Protocol

Use a versioned REST-style HTTP/JSON contract described with OpenAPI 3.1.

Initial contract path:

- /api/v1/capabilities

The initial contract is Development-only and reports server/API lifecycle and capability negotiation state. It does not make v1 Stable.

OpenAPI is stored in GoreeCloud/feeds-protocol as the cross-component authority. Server and clients must not silently redefine the contract.

### Shared implementation

Do not select a programming language or publish a package from GoreeCloud/feeds-shared yet.

Shared implementation begins only after at least two real consumers justify common ownership. Server-only and web-only implementation stays with its owning repository. Protocol definitions stay in feeds-protocol.

### Persistence

Database technology is intentionally deferred.

The roadmap strongly suggests durable relational data, but the first executable tranche does not require persistent user data. Persistence will be selected together with the first accepted core data model and migration design rather than prematurely coupling the project to a database.

### Authentication

Authentication/authorization implementation is intentionally deferred from the first executable tranche.

The target integration remains GoreeCloud Identity, with standards-based OIDC/OAuth 2.0 preferred where technically appropriate. No authentication claim is made until implemented and accepted.

### Deployment

Docker/container deployment is intentionally deferred until the server has persistent-state, configuration, health, recovery, and migration requirements that can be packaged truthfully.

## Consequences

Benefits:

- aligns each component with the preferred GoreeCloud language role;
- starts with low dependency burden;
- makes the cross-component contract explicit before feature proliferation;
- preserves technology independence by keeping protocol ownership separate from server code;
- avoids inventing shared abstractions before real reuse exists; and
- allows CI to validate exact candidate revisions from the first executable tranche.

Constraints:

- web UI work remains blocked until Glaze UI V1.6.0 integration is implemented and validated;
- protocol v1 remains Development, not Stable;
- database-connected persistence, authentication, scheduling, deployment, and production acceptance remain unresolved; PostgreSQL 18 is selected by ADR-0002 but is not yet a running dependency;
- changing a selected toolchain later requires a controlled decision and migration impact review.

## Verification sources

At decision time:

- Go 1.27.1 is the current stable Go patch release.
- Node.js 24.21.0 is an LTS release.
- TypeScript 7.0.2 is the latest stable npm release.
- Glaze UI V1.6.0 is the current Official Stable GoreeCloud consumer target.

External version choices should be reviewed during dependency/toolchain maintenance rather than treated as permanent architectural identities.
