# GoreeCloud Feeds Architecture

## Status

This document records the current Development architecture boundary for GoreeCloud Feeds. It does not establish production deployment, release, or Stable status.

## Product-family model

GoreeCloud Feeds is maintained as one product-family monorepo with distinct internal components:

- `services/server/` — authoritative server implementation.
- `apps/web/` — web client.
- `packages/protocol/` — versioned client/server contracts.
- `packages/shared/` — genuinely reusable implementation when reuse is demonstrated.

The repository boundary is shared because these components evolve as one product family. Their technical boundaries remain explicit.

## Server authority

The server is intended to be authoritative for:

- feed retrieval and scheduling;
- parsing and normalization;
- article processing, deduplication, and storage;
- search and rules;
- user/account state;
- notifications and media caching;
- synchronization and administration;
- feed health;
- backup and restore; and
- server API behavior.

Clients present reading, organization, search, settings, administration, and offline-capable interfaces while consuming shared versioned protocol definitions.

## Current Development implementation

The server source currently includes a Go runtime, normalized feed/article models, bounded RSS/Atom parsing, conservative source-scoped deduplication, PostgreSQL 18 migration/connectivity/repository foundations, and bounded remote-feed retrieval with SSRF-aware destination controls, redirect validation, conditional requests, concurrency limits, and transient retry/backoff.

The web source includes a pinned TypeScript/Node Development toolchain and typed capabilities client.

The protocol package contains the Development HTTP/JSON contract described with OpenAPI 3.1 and validation tooling.

The network-visible server runtime is not yet fully wired to all persistence and retrieval primitives. Production authentication, synchronization, complete user experience, deployment, and Stable qualification remain open.

## Persistence and compatibility

ADR-0001 selects Go for the server, TypeScript for the web client, and an HTTP/JSON Development contract. ADR-0002 selects PostgreSQL 18 as the Development persistence target with version-controlled SQL migrations.

Protocol definitions remain versioned so clients and server can evolve without unsupported compatibility claims.

## Future clients

Desktop, mobile, compact-reader, dashboard-widget, and administrative experiences may be added as additional app targets inside this product family unless a future boundary review establishes a genuinely independent lifecycle.
