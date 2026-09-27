# GoreeCloud Feeds Web

GoreeCloud Feeds Web is the Glaze UI web-client project for GoreeCloud Feeds.

## Current state

**Lifecycle:** Development.

The repository now contains a reproducible TypeScript Development toolchain, a typed Development-state module, and a dependency-free typed client for the non-sensitive `GET /api/v1/capabilities` Development contract. It does **not** yet contain a rendered web application, UI framework, Glaze UI implementation, authenticated product API client, offline runtime, browser bundle intended for users, deployment, Release Candidate, production acceptance, or Stable release.

## Development toolchain

- Node.js 24.21.0 LTS
- npm 11.19.0
- TypeScript 7.0.2
- strict TypeScript compiler configuration
- GitHub Actions type-check/build/protocol-client validation
- strict typed `0.1.0-dev` capability-response parsing and Development capability fetch client
- Development protocol target `0.1.0-dev`
- Glaze UI target `1.6.0`

The first tranche deliberately selects no UI framework. A framework may be introduced only when it materially improves accessibility, maintainability, testing, performance, and Glaze UI integration without unnecessary lock-in.

## Development validation

From the repository root with the pinned Node/npm toolchain:

```sh
npm install --ignore-scripts --no-audit --no-fund --package-lock=false
npm run typecheck
npm run build
npm run test:protocol
```

This produces Development modules and validates the capability client. It is not a user-facing application build or production deployment artifact.

## Planned responsibilities

The web client is intended to provide the GoreeCloud Feeds home dashboard, Today view, feed navigation, article lists and reader, unread management, saved articles, favorites, reading history, reading-position synchronization, search, folders, tags, Smart Feeds, rules management, settings, offline interface, connection/synchronization state, administration interface, responsive behavior, and accessibility.

## Repository relationships

- GoreeCloud/feeds — project-wide coordination and architecture decisions.
- GoreeCloud/feeds-server — authoritative service implementation.
- GoreeCloud/feeds-protocol — shared client/server contract authority.
- GoreeCloud/feeds-shared — reusable implementation only when genuine cross-repository reuse exists.

## Glaze UI

Live GoreeCloud GitHub state identifies **Glaze UI V1.6.0** as the current Official Stable consumer target. This repository records V1.6.0 as a required target, but no Glaze UI implementation or Feeds-specific conformance is yet verified.

## Documentation

See SPECIFICATIONS.md, FEATURES.md, FEATURE-ROADMAP.md, BENEFITS.md, COMPETITIVE-OBJECTIVES.md, BRANDING.md, USER-MANUAL.md, PRIVACY POLICY.md, SECURITY.md, NOTES.md, and CHANGELOG.md.

## Platform integration

The repository declares GoreeCloud Platform Contract schema 0.4 in goreecloud.platform.yaml. Application-specific integrations remain blocked until implementation and acceptance evidence exists.

## License

This repository currently uses the GoreeCloud fallback software license: GNU Affero General Public License v3.0 or later (AGPL-3.0-or-later), pending any later controlled project-specific licensing decision. See LICENSE.
