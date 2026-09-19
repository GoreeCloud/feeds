# GoreeCloud Feeds Architecture

## Status

This document records the **planned architecture boundary** for GoreeCloud Feeds. It does not establish implemented runtime behavior, deployment, release, or Stable status.

## Architectural model

GoreeCloud Feeds is planned as a client-server platform.

The server is intended to be authoritative for:

- feed retrieval and scheduling;
- feed parsing and normalization;
- article processing, deduplication, and storage;
- search;
- rules and Smart Feeds;
- user accounts;
- notifications;
- media caching;
- synchronization;
- administration;
- feed health;
- backup and restore; and
- the server API.

Clients are intended to present reading, organization, search, settings, administration, and offline-capable interfaces while communicating with the same server through shared protocol definitions.

## Initial repository boundaries

### `feeds`

Central project repository for project-wide architecture, development, deployment, contribution, roadmap, compatibility, release, and cross-repository coordination documentation.

### `feeds-server`

Planned complete back end and authoritative service implementation.

### `feeds-web`

Planned Glaze UI web client.

### `feeds-protocol`

Planned shared protocol contract for:

- API contracts;
- data structures;
- synchronization structures;
- error definitions;
- event definitions;
- versioning; and
- client/server compatibility rules.

### `feeds-shared`

Planned home for code that is genuinely reusable across more than one Feeds implementation boundary. Potential scope includes shared models, validation, feed utilities, common constants, shared synchronization logic, common transformations, and reusable test fixtures.

Repository-local convenience code should remain with the repository that owns it rather than being moved into `feeds-shared` prematurely.

## Authority and synchronization

The planned model keeps server-side feed and account state authoritative while clients maintain only the local state needed for interface behavior, caching, offline operation, and synchronization.

Protocol definitions should be versioned so clients and servers can evolve without unnecessarily breaking compatibility.

ADR-0001 now selects Go for the initial server runtime, TypeScript for the web toolchain, and a Development REST-style HTTP/JSON contract described with OpenAPI 3.1. The initial shared API path is `/api/v1/capabilities`. Database technology, queue/cache choices, web UI framework, authentication implementation, and deployment topology remain deferred.

## Future clients

The architecture is intended to allow additional clients, including desktop, mobile, compact-reader, dashboard-widget, and administrative experiences, to communicate with the same server contract.

Their repositories are not part of the initial implementation set unless separately authorized.
