# GoreeCloud Feeds Web Changelog

## Unreleased

### Development capability client

- Added strict TypeScript types and runtime validation for the authoritative `0.1.0-dev` capability response.
- Added a dependency-free Development fetch client for `GET /api/v1/capabilities`.
- Added HTTP(S)-only base-URL validation, embedded-credential rejection, credential omission, redirect rejection, no-referrer/no-store behavior, JSON media-type enforcement, and fail-closed version/schema checks.
- Added executable client validation to the existing exact-candidate CI workflow.
- Kept rendered UI, authentication, protected product APIs, offline state, deployment, and Stable qualification explicitly incomplete.


### TypeScript Development toolchain

- Selected Node.js 24.21.0 LTS, npm 11.19.0, and TypeScript 7.0.2 for the initial web Development toolchain.
- Added strict TypeScript compiler configuration.
- Added a typed Development-state module recording protocol target `0.1.0-dev`, Glaze UI target V1.6.0, and the absence of a rendered application.
- Added CI type-check/build validation.
- Kept UI framework selection, rendered Glaze UI implementation, server connection, authentication, offline behavior, deployment, and Stable qualification explicitly incomplete.

### Repository foundation

- Expanded the repository entry point.
- Added planned web-client specifications and scoped roadmap.
- Added mandatory repository documentation baseline.
- Added Platform Contract schema 0.4 declaration with application-specific integrations blocked pending implementation and acceptance.
- Recorded current Official Stable Glaze UI V1.6.0 as the future application adoption target without claiming conformance.
- Recorded the current GoreeCloud fallback license as AGPL-3.0-or-later.
- Preserved Planned lifecycle and avoided implementation, release, deployment, production, or Stable claims.
