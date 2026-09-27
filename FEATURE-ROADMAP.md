# GoreeCloud Feeds Feature Roadmap

## Status

**Lifecycle of this record:** planned capability roadmap.

**Implementation boundary:** controlled Development implementation is underway across the server, protocol, and web components inside this product-family monorepo. No capability is complete merely because source, schema, or a component foundation exists; each roadmap obligation still requires implementation and verification.

The authoritative GoreeCloud planning record is maintained in the governed Feature Roadmap location. This file is the repository-local roadmap representation for development coordination.

## Planned capability sections

1. Overview
2. Core Architecture
3. Feed Subscription Management
4. Feed Retrieval Engine
5. Feed Parsing and Normalization
6. Article Deduplication
7. Article Storage
8. Clean Article Reader
9. Optional Full Article Retrieval
10. Home Dashboard
11. Today View
12. Unread Management
13. Saved Articles
14. Favorites
15. Reading History
16. Reading Position Synchronization
17. Search
18. Smart Feeds
19. Rules Engine
20. Folders
21. Tags
22. Feed Health
23. Offline Mode
24. Connection State Labels
25. Multi-Device Synchronization
26. Notifications
27. Media Cache
28. Privacy
29. User Accounts
30. Shared Feeds and Collections
31. Administration
32. Backup and Restore
33. Glaze UI Integration
34. Accessibility
35. Responsive Interface
36. Core Internal Data Model
37. Client API
38. Future Native Clients
39. Repository Structure
40. Future Components
41. Product-family Development Model
42. Project Goal

## Product-family structure

The initial Feeds source family is consolidated into:

- `services/server/`
- `apps/web/`
- `packages/protocol/`
- `packages/shared/`

Future desktop, mobile, documentation-site, extension, or other clients should normally be added as component targets within this repository unless a later repository-boundary review establishes a genuine independent lifecycle.

## Current Development foundation

Verified Development source includes the Go/TypeScript/OpenAPI toolchain, PostgreSQL 18 persistence architecture, bounded parser/model/deduplication implementation, durable PostgreSQL migration/repository foundations, bounded remote retrieval with retry/backoff, a Development capability contract, and a typed web capability client.

Full runtime orchestration, authentication, synchronization, rendered Glaze UI product experience, production deployment, backup/restore acceptance, and Stable qualification remain incomplete.

## Status discipline

Each capability remains planned until implementation and authoritative verification demonstrate a stronger state.

Repository-level implementation work should update this roadmap when capability scope, priority, dependency, implementation state, verification state, or lifecycle disposition materially changes.
