# ADR-0002 — Initial GoreeCloud Feeds Persistence Architecture

## Status

Accepted for the Development persistence foundation.

This decision selects the initial durable datastore and migration boundary. It does not establish a deployed database, production data, database credentials, a Stable schema, a database driver, a backup implementation, or Release Candidate/Stable status.

## Decision date

2026-09-19

## Context

The first Development tranches established:

- the server runtime and versioned API foundation;
- normalized User, Subscription, Feed, Article, and ArticleState models;
- RSS 2.x and Atom 1.x normalization;
- conservative source-scoped article deduplication; and
- a bounded web capability client.

The governed roadmap now requires durable article retention, user-specific article state, preservation of selected articles, content fingerprints/source history, and eventually backup/restore of database contents.

GoreeCloud dependency governance requires significant databases to have explicit ownership, compatibility, security, backup, recovery, and failure behavior. GoreeCloud platform-contract governance requires persistent datasets to define schema, migration behavior, versioning, deletion, backup, restore, and portability.

## Decision

### Database

Use PostgreSQL 18 as the initial GoreeCloud Feeds durable relational datastore.

Development qualification baseline at this decision:

- PostgreSQL major version: 18;
- current supported minor verified at decision time: 18.6;
- production qualification, when it exists, must use a supported PostgreSQL 18.x minor and must not intentionally remain on an older vulnerable minor merely to preserve an arbitrary patch pin.

PostgreSQL is a required dependency only once a Feeds runtime actually enables durable persistence. Until a database-connected runtime is implemented and accepted, the dependency remains a Development architecture target rather than an operational requirement.

### Ownership

GoreeCloud Feeds Server owns the Feeds database and its schema.

Clients and other GoreeCloud applications must not access the Feeds database directly. Cross-component access uses supported Feeds APIs/protocols.

Database names, hostnames, ports, credentials, storage paths, and deployment topology remain deployment concerns and are not hard-coded by this architecture decision.

### Schema and identifiers

Version-controlled PostgreSQL migrations live in `GoreeCloud/feeds-server` under a conventional migration directory.

Initial schema principles:

- opaque application-owned textual identifiers rather than database row-order identities;
- shared Feed and Article content remains separate from user-owned Subscription and ArticleState records;
- database foreign keys and constraints protect referential integrity;
- saved/preserved user state remains independent from ordinary article-retention decisions;
- deduplication aliases are stored separately from canonical article content so publisher identifier/URL changes do not require rewriting article identity;
- timestamps use timezone-aware PostgreSQL timestamps;
- schema-visible state is explicit rather than hidden in undocumented serialized blobs unless a later requirement justifies one.

### Migration lifecycle

Migrations are:

- ordered;
- immutable after accepted application to a released/deployed environment;
- traceable to source control;
- reviewed with schema/data migration impact;
- validated before any future production promotion.

The first Development migration is an additive foundation. Destructive migration behavior is not introduced by ADR-0002.

### Database access

A PostgreSQL Go driver is intentionally deferred to the database-connectivity tranche.

When introduced, the driver must be:

- mature and maintained;
- license-compatible;
- pinned/locked as appropriate;
- documented in dependency records;
- used through parameterized queries or another safe typed/query boundary;
- replaceable without turning driver-specific behavior into the Feeds public protocol.

The application database identity must use least privilege and must not require PostgreSQL superuser rights for ordinary runtime operation.

### Backup and recovery

Database persistence makes database-aware backup and recovery mandatory before release qualification.

Future accepted backup/recovery must cover:

- database contents;
- migration/schema version;
- server configuration required to reconstruct the service;
- any separately stored media/cache state that is non-reconstructable;
- integrity verification;
- isolated restore validation.

Everkeep remains the required GoreeCloud recovery integration before production acceptance where applicable. A schema or migration file alone is not backup/restore implementation evidence.

### Search, cache, and queues

ADR-0002 does not select:

- a separate search service;
- Redis/Valkey;
- a message queue;
- a background-job system;
- PostgreSQL extensions beyond core PostgreSQL;
- a container image or deployment topology.

These dependencies must be justified separately if later requirements make them necessary.

## Consequences

Benefits:

- relational constraints fit the current Feeds ownership model;
- PostgreSQL provides a mature self-hostable and portable datastore;
- plain version-controlled SQL keeps the data contract understandable and migration history inspectable;
- separating canonical content, user state, and deduplication aliases supports multi-user isolation and publisher metadata changes;
- the design does not add a cache, queue, or search dependency prematurely.

Constraints:

- persistence is not usable until a database connectivity/transaction layer exists;
- schema migrations become compatibility/recovery obligations once deployed;
- future user/account persistence must integrate with GoreeCloud Identity without storing authentication secrets in Feeds;
- backup and restore become release blockers once unique durable data exists;
- data retention/deletion behavior must preserve explicitly retained user content.

## Verification sources

At decision time:

- GoreeCloud roadmap sections 7, 32, and 36 require article storage, database-aware backup/restore scope, and separation of shared content from user-specific state.
- GoreeCloud dependency governance explicitly recognizes PostgreSQL as an application database and requires documented ownership, version compatibility, backup, recovery, authentication, and network relationships.
- GoreeCloud programming strategy identifies SQL as the relational schema/query language and prefers PostgreSQL-compatible SQL where PostgreSQL is selected.
- PostgreSQL project versioning information identifies 18.6 as the current supported PostgreSQL 18 minor on 2026-09-19.

External product-version state must be revalidated during dependency maintenance and release qualification rather than treated as permanent.
