# GoreeCloud Feeds Web Specifications

## Status

Component: GoreeCloud Feeds Web  
Repository: GoreeCloud/feeds  
Component class: Web application  
Lifecycle: Development  
Implementation status: TypeScript Development toolchain plus strict capability and bounded article-list protocol clients established; rendered application not yet implemented

This specification scopes the web-client responsibilities derived from the governing GoreeCloud Feeds product roadmap. The current implementation includes a pinned TypeScript Development toolchain, typed Development-state module, a strict client for non-sensitive `GET /api/v1/capabilities`, and a strict bounded client for `GET /api/v1/articles`; no rendered user-facing application is implemented.

## Authority boundary

GoreeCloud Feeds Web is intended to present and manage authorized GoreeCloud Feeds state while GoreeCloud Feeds Server remains authoritative for server-owned feed ingestion, processing, search, storage, synchronization, and administration state.

The web client must not silently create an incompatible second authority for server-owned data.

## Planned interface capabilities

- Home dashboard and Today view.
- Feed navigation and article lists.
- Clean article reader.
- Unread management.
- Saved articles and favorites.
- Reading history.
- Reading-position synchronization.
- Search.
- Folders and tags.
- Smart Feeds.
- Rules management.
- Settings.
- Feed-health presentation.
- Offline interface where safe and practical.
- Explicit connection and synchronization state.
- Notification-related settings/presentation where supported.
- Authorized administration interface.
- Responsive desktop/tablet/mobile presentation.
- Keyboard and touch-friendly interaction.
- Accessibility and reduced-motion/transparency behavior.
- Current Official Stable Glaze UI adoption.

## Protocol boundary

Client/server communication should use versioned shared contracts owned by GoreeCloud/feeds-protocol. The web client should not depend on undocumented server internals.

## Offline boundary

Offline behavior should preserve local usability and user changes where supported without silently overwriting server-authoritative state. Synchronization failures and conflicts must remain visible and recoverable.

## Privacy and security

The client should minimize unnecessary external requests, avoid advertising/profiling dependencies, protect session and account state, and avoid leaking article/feed content through diagnostics.

## Current implementation decision

The web language is TypeScript 7.0.2. Development/CI execution uses Node.js 24.21.0 LTS and npm 11.19.0. The repository currently uses the TypeScript compiler directly and intentionally has no UI framework or browser bundler. `src/protocol/capabilities.ts` implements strict `0.1.0-dev` capability-response validation and a dependency-free fetch client for non-sensitive capability discovery. `src/protocol/articles.ts` implements strict bounded article-list validation and a credential-aware request path for server-derived user context. Both clients reject redirects, use no-store behavior, and validate protocol shape before returning data. Glaze UI V1.6.0 remains mandatory for the first rendered GoreeCloud interface.

## Open decisions

UI framework, browser bundling/application packaging, authenticated product API client design, authentication/session handling, offline-storage engine, service-worker design, browser support matrix, end-to-end test tooling, deployment host, and production packaging remain unresolved.
