# GoreeCloud Feeds Protocol Specifications

## Status

Component: GoreeCloud Feeds Protocol  
Repository: GoreeCloud/feeds-protocol  
Component class: Shared protocol / contract repository  
Lifecycle: Planned  
Implementation status: Repository documentation foundation only

This specification defines the ownership boundary for future shared protocol work. It does not establish an implemented or released protocol.

## Protocol authority

GoreeCloud/feeds-protocol is intended to be the authoritative repository for client/server contracts that must be understood by more than one GoreeCloud Feeds component.

Server-private persistence schemas, internal queue formats, private implementation details, browser-only UI state, and repository-local helper types do not become protocol contracts merely because they exist.

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

No serialization format, transport, schema language, code-generation tool, package format, versioning syntax, API base path, event transport, or compatibility window is selected by this foundation.
