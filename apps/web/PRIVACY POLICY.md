# GoreeCloud Feeds Web Privacy Policy

## Current implementation boundary

This repository does not yet contain a working web application. It does contain a Development-only capability client that requests non-sensitive server lifecycle/capability metadata. The broader statements below define planned privacy requirements and are not claims about deployed product behavior.

## Current Development client behavior

The capability client sends no user content, omits credentials, sends no referrer, requests no-store behavior, rejects redirects, and processes only the `0.1.0-dev` capability response. It does not use cookies, sessions, browser-local durable state, telemetry, or third-party analytics.

## Planned privacy requirements

The web client should avoid advertising, profiling, and mandatory third-party analytics; communicate with the configured GoreeCloud Feeds Server through documented interfaces; minimize external requests; avoid leaking feed/article content through diagnostics; keep offline/browser-local data limited to what supported functionality requires; make local-cache deletion/clearing behavior understandable; and integrate with GoreeCloud Privacy Shield before accepted production qualification where applicable.

Future cookies, tokens, browser storage, cache scope, external endpoints, telemetry, diagnostics, retention, and account/session behavior must be documented from verified implementation before release.
