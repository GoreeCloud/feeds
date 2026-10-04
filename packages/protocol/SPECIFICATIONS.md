# GoreeCloud Feeds Protocol Specifications

## Status

Component: GoreeCloud Feeds Protocol  
Repository: GoreeCloud/feeds  
Component class: Shared protocol / contract repository  
Lifecycle: Development  
Implementation status: Versioned OpenAPI 3.1 Development contract and validation tooling implemented for capability discovery and bounded article listing

This specification defines the ownership boundary for shared Feeds client/server contracts. The current `0.1.0-dev` OpenAPI contract is source-implemented and validated, but it is Development-only and does not establish a Stable or production protocol release.

## Protocol authority

GoreeCloud/feeds-protocol is intended to be the authoritative repository for client/server contracts that must be understood by more than one GoreeCloud Feeds component.

Server-private persistence schemas, internal queue formats, private implementation details, browser-only UI state, and repository-local helper types do not become protocol contracts merely because they exist.

## Current Development contract

The verified Development contract currently includes:

- `GET /api/v1/capabilities` for non-sensitive protocol and capability negotiation.
- `GET /api/v1/articles` for a bounded chronological article list scoped to server-derived authenticated or approved local user context.
- `articles:list-v1` as the capability identifier for the article-list feature.
- An optional `limit` parameter bounded to 1..100 with a default of 50.
- Article summary fields for article/feed identity, feed title, URL, title, author, publication time, summary, language, and read/saved/favorite state.
- Explicit 400, 401, and 503 error states for invalid requests, unavailable identity, or unavailable backing services.
- No client-supplied user identifier in the article-list contract.

The contract does not itself define the production authentication/session or local-only identity mechanism. The server must derive that context through an approved implementation before advertising article listing.

## Planned contract areas

### API contracts

Shared request, response, resource, pagination, filtering, authentication-context, and error structures should be defined here when they become cross-component interfaces.

### Data structures

Only data structures that are part of supported client/server or event interfaces should be protocol-owned.

### Synchronization

Synchronization contracts should define authoritative state, revision/version fields, incremental changes, retries, interrupted-sync recovery, deletions/tombstones, conflicts, and schema evolution where applicable.

### Errors

Errors intended for clients should have stable machine-readable identities and documented semantics. Internal stack traces or server-private exceptions must not become accidental public contracts.

### Events

Cross-component events should define names, versions, payloads, ordering assumptions, delivery semantics, authorization/privacy boundaries, and compatibility expectations where applicable.

### Capability negotiation

Clients and servers should be able to determine supported optional behavior without depending on undocumented implementation details when version differences require negotiation.

## Security and privacy

Protocols must not require reusable secrets in payloads or expose private implementation data unnecessarily. Authentication and authorization claims must be backed by the owning security/identity systems.

## Open decisions

Production authentication/session semantics, local-only identity semantics, synchronization contracts, write-state APIs, search/source/folder APIs, event transport, code generation, package format, and the Stable compatibility window remain open. HTTP/JSON, OpenAPI 3.1, `/api/v1/`, and `0.1.0-dev` are selected for the current Development contract.
