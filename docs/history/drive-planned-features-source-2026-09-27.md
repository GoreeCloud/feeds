# Historical Drive Planned Features Migration Source — GoreeCloud Feeds

> **Status:** Historical, non-authoritative migration evidence.  
> **Source:** Former Google Drive planning document, captured during repository migration on 2026-09-27.  
> **Rule:** Do not synchronize this file with Google Drive. Current feature truth is in `IMPLEMENTED-FEATURES.md`, `PLANNED-FEATURES.md`, and `CHANGELOGS.md`.

GOREECLOUD
Feeds
Planned Features and Capabilities
Contents
Section index for the complete planning specification.
1. Overview
GoreeCloud Feeds is a fully self-hosted feed aggregation, synchronization, reading, discovery, and content-management platform built as part of the GoreeCloud ecosystem.
The project will include both a dedicated server-side back end and a Glaze UI-based front end. The server will be responsible for discovering, retrieving, processing, normalizing, storing, indexing, synchronizing, and serving feed content. The client will provide a polished interface for browsing, reading, searching, organizing, and managing that content.
GoreeCloud Feeds should be designed so that the server remains the authoritative source of feed data while web, desktop, mobile, and future GoreeCloud clients can connect to the same service.
The platform should prioritize:
Self-hosting
Privacy
Local data ownership
Offline access
Fast synchronization
Accessibility
Multi-device support
Extensibility
Glaze UI integration
Minimal external dependencies
2. Core Architecture
GoreeCloud Feeds will use a client-server architecture.
GoreeCloud Feeds Server
The server will handle:
Feed subscriptions
Feed retrieval
Feed parsing
Content normalization
Article storage
Search indexing
User accounts
Read/unread state
Favorites and saved articles
Tags
Folders
Smart Feeds
Rules
Notifications
Offline synchronization
Content caching
Feed health monitoring
Administrative controls
Client synchronization
GoreeCloud Feeds Client
The client will provide:
Feed browsing
Article reading
Feed management
Folder management
Tag management
Search
Smart Feeds
Saved articles
Favorites
Reading history
Offline access
Server synchronization
User preferences
Notifications
Administrative interfaces when authorized
3. Feed Subscription Management
Users will be able to subscribe directly to feeds and organize them however they prefer.
Planned capabilities include:
Subscribe using a feed URL
Automatic feed discovery from supported websites
Manual feed configuration
Feed import
Feed export
Subscription backup
Subscription restoration
Feed renaming
Custom feed descriptions
Custom feed icons
Feed categories
Folders
Nested folders
Tags
Favorite feeds
Muted feeds
Archived feeds
Hidden feeds
Per-feed refresh settings
Per-feed retention settings
Per-feed notification settings
Feed health information
Feed subscription history
Multiple users should be able to subscribe to the same source while maintaining independent personal settings and reading states.
4. Feed Retrieval Engine
The server will operate its own feed retrieval system.
The retrieval engine should provide:
Scheduled feed refreshes
Manual refresh
Adaptive refresh scheduling
Priority feed refreshes
Failed request retrying
Backoff behavior
Retrieval queue management
Request timeout handling
Connection error handling
Feed change detection
Conditional retrieval where supported
Bandwidth-aware retrieval
Parallel feed processing
Configurable concurrency
Fetch history
Last successful retrieval tracking
Frequently updated feeds may be refreshed more often than feeds that rarely change.
5. Feed Parsing and Normalization
Different feed formats should be converted into a common internal GoreeCloud Feeds representation.
The parser should extract and normalize:
Feed title
Feed description
Feed icon
Feed image
Article title
Article identifier
Article URL
Author
Publication date
Updated date
Summary
Full feed content
Categories
Tags
Embedded media
Images
Language
Source information
Malformed or incomplete feeds should be handled gracefully whenever usable information can still be recovered.
6. Article Deduplication
GoreeCloud Feeds should prevent duplicate articles from appearing when publishers modify identifiers, URLs, or metadata.
Deduplication may use combinations of:
Source identifiers
Canonical URLs
Normalized URLs
Article titles
Publication timestamps
Authors
Content fingerprints
Content similarity
Feed source relationships
Potential duplicates should be detected without incorrectly combining unrelated articles.
7. Article Storage
Articles should remain available according to configurable retention rules instead of disappearing simply because they are no longer present in the source feed.
Stored article information may include:
Original feed content
Normalized article content
Article metadata
Media references
Publication information
Retrieval timestamps
Content fingerprints
Source history
User-specific state
Users should be able to permanently preserve selected articles regardless of ordinary retention policies.
8. Clean Article Reader
GoreeCloud Feeds should provide multiple reading modes.
Original View
Displays the content supplied directly through the feed.
Clean View
Displays normalized article content in a cleaner reading layout.
Focus View
Provides a distraction-reduced reading environment with minimal interface elements.
Reader settings should include:
Font size
Reading width
Line spacing
Paragraph spacing
Content density
Image visibility
Reading theme
Reader zoom
Text alignment
Per-user defaults
9. Optional Full Article Retrieval
Some feeds provide only article excerpts.
GoreeCloud Feeds may optionally retrieve additional readable article content where permitted and available.
The system should retain both:
Original feed content
Retrieved article content
Users should always be able to distinguish between the two.
Full-content retrieval should be configurable globally and per feed.
10. Home Dashboard
The Home interface should provide an overview of the user's reading environment.
Possible dashboard sections include:
Recently published
Recently updated
Unread articles
Saved articles
Favorite feeds
Priority feeds
Smart Feeds
Continue reading
Recently read
Feed health
Synchronization status
Users should eventually be able to customize which dashboard sections are displayed.
11. Today View
The Today view will provide a chronological summary of recently published content.
Users should be able to filter Today by:
All feeds
Folder
Tag
Favorite feeds
Smart Feed
Read state
Publication time
Content type
12. Unread Management
Unread content will be treated as a primary organizational state.
Capabilities should include:
Mark article as read
Mark article as unread
Mark feed as read
Mark folder as read
Mark articles above as read
Mark articles below as read
Mark articles older than a selected date as read
Automatically mark articles as read
Read-state synchronization
Unread counts
Folder unread counts
Feed unread counts
13. Saved Articles
Users will be able to save articles for later.
Saved articles should remain accessible independently from their original feed organization.
Capabilities should include:
Save article
Remove from saved
Saved article search
Saved article tags
Saved article notes
Permanent preservation
Offline availability
Saved article filtering
14. Favorites
Users should be able to favorite both articles and feeds.
Favorite content may receive:
Dedicated navigation
Priority synchronization
Priority notification options
Custom filtering
Permanent preservation options
15. Reading History
GoreeCloud Feeds should maintain optional reading history.
History information may include:
Article opened
Article completed
Last opened time
Reading position
Time spent reading
Device synchronization
Users should be able to clear or disable history according to their preferences.
16. Reading Position Synchronization
Long articles should support synchronized reading progress.
The server may store:
Last reading position
Completion status
Last opened timestamp
Opening the same article on another authorized client should allow the user to continue from approximately the same location.
17. Search
GoreeCloud Feeds should provide fast full-text search.
Searchable fields should include:
Article title
Article body
Summary
Author
Feed name
Folder
Tags
Notes
Saved content
Search filters should include:
Date range
Feed
Folder
Tag
Author
Read state
Saved state
Favorite state
18. Smart Feeds
Smart Feeds will provide automatically generated collections based on rules.
Rules may include:
Feed
Folder
Tag
Author
Keyword
Title
Article content
Date
Read state
Saved state
Favorite state
Rules should support combinations such as:
Unread articles containing selected keywords from selected folders published during the last seven days.
Smart Feeds should update automatically as new content arrives.
19. Rules Engine
A general-purpose rules engine should allow automated content organization.
Rules may perform actions such as:
Add tag
Remove tag
Mark as read
Save article
Favorite article
Hide article
Move into a Smart Feed
Trigger a notification
Increase article priority
Mute content
Rules should support multiple conditions and actions.
20. Folders
Subscriptions should support hierarchical organization.
Example structure:
Technology
→ Development
→ Infrastructure
→ Security
Capabilities should include:
Create folder
Rename folder
Delete folder
Move folder
Nest folder
Move feeds between folders
Folder unread counts
Folder-level settings
Folder-level notification settings
21. Tags
Tags should provide organization independent of folders.
Tags may be applied to:
Articles
Feeds
Saved articles
Tags should support:
Custom names
Glaze UI semantic colors
Search
Filtering
Automatic rules
22. Feed Health
GoreeCloud Feeds should include a dedicated feed health system.
Possible states include:
Healthy
Delayed
Changed
Unavailable
Authentication Required
Invalid
Needs Attention
Feed health information may show:
Last successful retrieval
Last attempted retrieval
Last new article
Response status
Parsing status
Feed URL changes
Consecutive failures
The system should preserve previously retrieved content when a feed becomes unavailable.
23. Offline Mode
The client should remain useful without a connection to the GoreeCloud Feeds Server.
Offline capabilities should include:
Cached article reading
Cached feed browsing
Saved articles
Previously synchronized search data
Queued read/unread changes
Queued favorites
Queued saved-state changes
Queued tag changes
Queued notes
Changes should synchronize automatically after connectivity returns.
24. Connection State Labels
The interface should clearly communicate where content is coming from.
Possible states include:
Online
Offline
Cached
Synchronizing
Server Unavailable
Updating
Out of Date
Status should not rely on color alone.
Icons, text labels, and accessibility information should also communicate the state.
25. Multi-Device Synchronization
The server should synchronize user state between authorized GoreeCloud Feeds clients.
Synchronized information should include:
Subscriptions
Folders
Tags
Read state
Saved state
Favorites
Notes
Reading history
Reading position
Smart Feeds
Rules
Preferences
Notification settings
26. Notifications
Users should be able to create notifications for important feed activity.
Notification conditions may include:
New article from selected feed
New article from selected folder
Keyword match
Author match
Smart Feed match
Priority article
Feed health failure
Notifications should be configurable per user.
27. Media Cache
The server may optionally cache article media locally.
Possible cached content includes:
Article images
Feed icons
Feed artwork
Other supported article assets
Cache management should support:
Storage limits
Automatic cleanup
Per-user limits
Per-feed settings
Offline availability
Cache inspection
28. Privacy
GoreeCloud Feeds should be designed around private self-hosting.
Privacy goals include:
Subscription data remains on the GoreeCloud infrastructure selected by the administrator
Reading history remains private
Saved articles remain private
Search queries remain private
No outside behavioral analytics
No advertising tracking
No external profiling
No unnecessary remote requests from clients
Administrative access should not automatically expose user reading activity
Privacy-sensitive features should be configurable by deployment administrators and users where appropriate.
29. User Accounts
The server should support multiple users.
Each user should have independent:
Subscriptions
Folders
Tags
Read state
Saved articles
Favorites
History
Smart Feeds
Rules
Preferences
Notifications
Shared content should require explicit configuration.
30. Shared Feeds and Collections
A future collaboration system may allow users to share selected feed collections.
Possible shared objects include:
Shared folders
Shared feed lists
Curated collections
Shared Smart Feeds
Personal reading activity should remain separate unless explicitly shared.
31. Administration
Administrators should receive a dedicated management interface.
Administration capabilities should include:
Server status
User management
Storage usage
Feed count
Article count
Retrieval queue
Failed feeds
Parser failures
Synchronization status
Cache usage
Database health
Retention settings
Backup management
Restore management
Audit history
Server configuration
32. Backup and Restore
GoreeCloud Feeds should provide built-in backup support.
Backups may include:
Server configuration
Users
Subscriptions
Folders
Tags
Smart Feeds
Rules
Saved articles
Article state
Database contents
Administrators should be able to verify backup integrity before restoration.
33. Glaze UI Integration
The GoreeCloud Feeds front end should use Glaze UI as its design system.
The interface should make extensive use of:
Semantic colors
Transparency
Translucency
Blur
Layered surfaces
Opacity
Depth
Smooth transitions
Motion
Skeleton loading states
Responsive layouts
Contextual status indicators
Adaptive navigation
Consistent spacing
Consistent typography
The interface should feel polished, modern, visually rich, and distinctly GoreeCloud.
34. Accessibility
Accessibility should be treated as a core capability.
Planned accessibility features include:
Keyboard navigation
Screen-reader support
Visible focus states
Semantic interface structure
Accessible labels
Reduced-motion support
High-contrast support
Scalable text
Large interaction targets
Status indicators that do not depend solely on color
35. Responsive Interface
The web client should adapt to:
Desktop
Laptop
Tablet
Mobile
Narrow windows
Wide displays
Desktop layouts may use:
Navigation → Article List → Article Reader
Smaller displays may collapse these into individual navigable views.
36. Core Internal Data Model
The platform should use an internal model similar to:
User → Subscription → Feed → Article → Article State
Additional entities may include:
Folder
Tag
Smart Feed
Rule
Note
Notification
Feed Health Record
Retrieval Record
Client Device
Synchronization Record
Feed content and user-specific article state should remain separated.
This allows many users to consume the same feed while maintaining completely independent states.
37. Client API
The GoreeCloud Feeds Server should expose a stable first-party interface for GoreeCloud clients.
The interface should support:
Authentication
Subscriptions
Feeds
Articles
Folders
Tags
Search
Smart Feeds
Rules
Saved articles
Favorites
Reading state
Reading position
Notifications
Preferences
Synchronization
Administration
The protocol should be versioned so clients and servers can evolve without unnecessarily breaking compatibility.
38. Future Native Clients
The server architecture should allow additional GoreeCloud Feeds clients to be developed later.
Potential clients include:
GoreeCloud Feeds Web
GoreeCloud Feeds Desktop
GoreeCloud Feeds Mobile
Compact sidebar reader
Dashboard widgets
Administrative client
All clients should communicate with the same GoreeCloud Feeds Server.
39. Repository Structure
The project should be split into focused repositories rather than placing the entire platform in one repository.
Required Repositories
feeds
Main project repository.
Responsibilities:
Project overview
Architecture documentation
Development documentation
Deployment documentation
Contribution guidelines
Project roadmap
Cross-repository coordination
Compatibility documentation
Release documentation
This repository acts as the central home for the GoreeCloud Feeds project.
feeds-server
Contains the complete GoreeCloud Feeds back end.
Responsibilities:
Feed retrieval
Feed scheduling
Parsing
Normalization
Article processing
Article storage
Search
Deduplication
Rules engine
Smart Feeds
Notifications
Media caching
User accounts
Synchronization
Administration
Feed health
Backup and restore
Server API
feeds-web
Contains the Glaze UI web client.
Responsibilities:
Home dashboard
Feed navigation
Article lists
Article reader
Search
Folders
Tags
Saved articles
Favorites
History
Smart Feeds
Rules management
Settings
Offline interface
Synchronization interface
Administration interface
Responsive design
Accessibility
feeds-protocol
Contains the shared GoreeCloud Feeds protocol definitions.
Responsibilities:
API contracts
Data structures
Synchronization structures
Error definitions
Event definitions
Versioning
Client/server compatibility rules
Keeping the protocol separate allows future clients to use the same server contract.
feeds-shared
Contains reusable GoreeCloud Feeds components that are not specific to the server or a particular client.
Potential responsibilities:
Shared models
Validation
Feed-related utilities
Common constants
Shared synchronization logic
Common data transformations
Shared testing fixtures
Only genuinely reusable code should be placed here.
40. Future Repositories
These repositories do not need to be created during the initial development phase but should be reserved for future expansion.
goreecloud-feeds-desktop
Native GoreeCloud Feeds desktop client.
goreecloud-feeds-mobile
Native GoreeCloud Feeds mobile client.
goreecloud-feeds-docs
Dedicated documentation site if the documentation eventually becomes too large for the main project repository.
goreecloud-feeds-extensions
Optional extensions and first-party integrations for the GoreeCloud Feeds ecosystem.
41. Recommended Initial Repository Set
Development should begin with five repositories:
feeds
feeds-server
feeds-web
feeds-protocol
feeds-shared
This creates a clear separation between project governance, server development, client development, shared protocol definitions, and reusable internal code.
The desktop and mobile repositories can be created when development of those clients begins.
42. Project Goal
The long-term goal of GoreeCloud Feeds is to become the GoreeCloud ecosystem's central platform for consuming, organizing, searching, synchronizing, preserving, and reading syndicated content.
Rather than functioning as only a basic feed reader, GoreeCloud Feeds should operate as a complete personal information stream with:
GoreeCloud Feeds Server
as the authoritative feed ingestion, processing, search, storage, and synchronization service,
and
GoreeCloud Feeds
as the polished Glaze UI reading and management experience presented to the user.
The architecture should make GoreeCloud Feeds useful as a standalone application while also allowing its feed, article, search, notification, and synchronization capabilities to be integrated into other GoreeCloud applications in the future.
PLANNING RECORD - This document defines intended capabilities. It does not establish implementation, deployment, release, or Stable status.
PRODUCT
GoreeCloud Feeds
DOCUMENT TYPE
Product Feature Roadmap / Capability Specification
STATUS
Planned
VERSION
v0.1
CLASSIFICATION
Internal
CREATED
September 19, 2026
AUTHORITATIVE SCOPE
Planned GoreeCloud Feeds product capabilities and architecture
SOURCE
Owner-submitted planning specification
1. Overview
22. Feed Health
2. Core Architecture
23. Offline Mode
3. Feed Subscription Management
24. Connection State Labels
4. Feed Retrieval Engine
25. Multi-Device Synchronization
5. Feed Parsing and Normalization
26. Notifications
6. Article Deduplication
27. Media Cache
7. Article Storage
28. Privacy
8. Clean Article Reader
29. User Accounts
9. Optional Full Article Retrieval
30. Shared Feeds and Collections
10. Home Dashboard
31. Administration
11. Today View
32. Backup and Restore
12. Unread Management
33. Glaze UI Integration
13. Saved Articles
34. Accessibility
14. Favorites
35. Responsive Interface
15. Reading History
36. Core Internal Data Model
16. Reading Position Synchronization
37. Client API
17. Search
38. Future Native Clients
18. Smart Feeds
39. Repository Structure
19. Rules Engine
40. Future Repositories
20. Folders
41. Recommended Initial Repository Set
21. Tags
42. Project Goal