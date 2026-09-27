# GoreeCloud Feeds Web Features

## Current implemented state

Verified Development foundations:

- pinned Node.js 24.21.0 LTS / npm 11.19.0 toolchain requirement;
- TypeScript 7.0.2;
- strict TypeScript compiler configuration;
- typed Development-state module recording product, lifecycle, protocol target, Glaze UI target, and absence of a rendered application;
- automated TypeScript type-check and build workflow;
- strict typed validation for the `0.1.0-dev` `GET /api/v1/capabilities` response;
- dependency-free capability fetch client with HTTP(S)-only URL validation, credential omission, redirect rejection, no-referrer/no-store behavior, JSON content-type enforcement, and fail-closed protocol checks;
- executable Development protocol-client validation.

No user-facing web capability is currently implemented. The repository can issue the non-sensitive Development capability request, but there is no rendered interface, authenticated product API integration, authentication flow, offline runtime, or deployable web application.

## Planned capability groups

- Home dashboard
- Today view
- Feed navigation
- Article lists and clean reader
- Unread management
- Saved articles
- Favorites
- Reading history
- Reading-position synchronization
- Search
- Folders and tags
- Smart Feeds
- Rules management
- Settings
- Feed-health presentation
- Offline interface
- Connection and synchronization state
- Administration interface
- Responsive layout
- Accessibility
- Glaze UI V1.6.0 adoption

Planned capability text must not be interpreted as current functionality.
