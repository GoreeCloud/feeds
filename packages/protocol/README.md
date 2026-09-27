# GoreeCloud Feeds Protocol

GoreeCloud Feeds Protocol is the planned shared contract repository for communication between GoreeCloud Feeds Server and current or future GoreeCloud Feeds clients.

## Current state

**Lifecycle:** Planned / repository foundation.

This repository now contains its first machine-readable **Development** API contract in `openapi/feeds-v1.json`. It does not define a Stable protocol, synchronization schema, event transport, generated SDK, compatibility matrix, or production release.

## Planned responsibilities

This repository is intended to own shared GoreeCloud Feeds protocol definitions for:

- API contracts;
- shared data structures;
- synchronization structures;
- error definitions;
- event definitions;
- versioning;
- capability negotiation where required;
- compatibility and deprecation rules; and
- client/server migration boundaries.

The protocol repository must describe cross-component contracts, not server-private implementation details.

## Repository relationships

- GoreeCloud/feeds — project-wide coordination and roadmap.
- GoreeCloud/feeds-server — planned authoritative server implementation.
- GoreeCloud/feeds-web — planned web client.
- GoreeCloud/feeds-shared — reusable implementation that is genuinely shared but is not itself a protocol authority.

## Development contract

The first contract tranche defines `GET /api/v1/capabilities` using OpenAPI 3.1 and HTTP/JSON. Contract version `0.1.0-dev` is Development-only and must not be represented as Stable.

## Documentation

See SPECIFICATIONS.md, VERSIONING.md, COMPATIBILITY.md, SECURITY.md, CONTRIBUTING.md, NOTES.md, and CHANGELOG.md.

## License

This public repository currently uses the GoreeCloud fallback software license: GNU Affero General Public License v3.0 or later (AGPL-3.0-or-later), pending any later controlled project-specific licensing decision. See LICENSE.
