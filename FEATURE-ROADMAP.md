# GoreeCloud Feeds Feature Roadmap

## Status

**Lifecycle of this record:** planned capability roadmap.

**Implementation boundary:** controlled Development implementation is underway across the server, protocol, and web repositories. No capability is complete merely because code, schema, or repository foundation exists; each roadmap obligation still requires its own implementation and verification.

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
40. Future Repositories
41. Recommended Initial Repository Set
42. Project Goal

## Initial repository set

The initial repository set is established as:

- `feeds`
- `feeds-server`
- `feeds-web`
- `feeds-protocol`
- `feeds-shared`

Repository existence does not establish implementation completion.

## Deferred repository expansion

Desktop, mobile, dedicated documentation-site, and extension repositories remain deferred until their development phases are authorized.

## Status discipline

Each capability remains planned until implementation and authoritative verification in the owning repository demonstrate a stronger state.

Repository-level implementation work should update this roadmap when capability scope, priority, dependency, implementation state, verification state, or lifecycle disposition materially changes.
\n## Current Development foundation\n\nVerified project-level decisions now include the initial Go/TypeScript/OpenAPI toolchain and PostgreSQL 18 persistence architecture. The server has bounded parser/model/deduplication implementation, the protocol has a Development capability contract, and the web repository has a typed capability client. Durable database connectivity, storage runtime, backup/restore, remote retrieval, authentication, synchronization, rendered Glaze UI, deployment, and Stable qualification remain incomplete.\n