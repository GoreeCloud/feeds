# GoreeCloud Feeds Shared Specifications

## Status

Component: GoreeCloud Feeds Shared  
Repository: GoreeCloud/feeds-shared  
Component class: Shared library / reusable package repository  
Lifecycle: Planned  
Implementation status: Repository documentation foundation only

This specification defines the ownership boundary for future reusable implementation. It does not establish a package or public API.

## Ownership boundary

GoreeCloud/feeds-shared is intended for implementation that has demonstrated reuse across multiple GoreeCloud Feeds components.

A type, helper, utility, fixture, or model does not become shared merely because it could theoretically be reused later.

## Appropriate future scope

Potential shared scope includes:

- reusable internal models that are not protocol authority;
- common validation;
- feed-specific utility functions;
- shared constants;
- common normalization/transformation helpers;
- synchronization implementation helpers that do not define the protocol itself; and
- reusable testing fixtures.

## Excluded scope

The following should remain with their owning repositories unless a later verified refactor changes the ownership model:

- server-only ingestion, storage, scheduling, search, or administration logic;
- web-only view, browser, offline, routing, or presentation state;
- client/server API and synchronization contract definitions;
- platform-specific packaging or deployment logic; and
- speculative abstractions with only one real consumer.

## Open decisions

No programming language, package manager, module format, public API, dependency graph, source layout, test framework, package registry, versioning scheme, or release process is selected by this foundation.
