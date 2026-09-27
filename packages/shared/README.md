# GoreeCloud Feeds Shared

GoreeCloud Feeds Shared is the planned repository for reusable GoreeCloud Feeds implementation components that are genuinely shared across more than one Feeds repository.

## Current state

**Lifecycle:** Planned / repository foundation.

This repository currently contains documentation only. It does not yet contain a shared package, runtime library, published artifact, source modules, test suite, dependency manifest, generated code, or Stable release.

## Scope rule

Code belongs here only when real reuse across multiple GoreeCloud Feeds components justifies shared ownership.

Server-only code belongs in GoreeCloud/feeds-server. Web-only code belongs in GoreeCloud/feeds-web. Shared client/server contracts belong in GoreeCloud/feeds-protocol.

This repository must not become a miscellaneous dumping ground or a competing protocol authority.

## Potential responsibilities

When actual reuse is established, this repository may contain:

- shared models that are not themselves public protocol contracts;
- common validation;
- feed-related utilities;
- shared constants;
- reusable synchronization helpers;
- common data transformations; and
- reusable test fixtures.

## Repository relationships

- GoreeCloud/feeds — project-wide coordination and roadmap.
- GoreeCloud/feeds-server — server-specific implementation.
- GoreeCloud/feeds-web — web-specific implementation.
- GoreeCloud/feeds-protocol — shared client/server protocol authority.

## Documentation

See SPECIFICATIONS.md, REUSE-GUIDELINES.md, COMPATIBILITY.md, SECURITY.md, CONTRIBUTING.md, NOTES.md, and CHANGELOG.md.

## License

This public repository currently uses the GoreeCloud fallback software license: GNU Affero General Public License v3.0 or later (AGPL-3.0-or-later), pending any later controlled project-specific licensing decision. See LICENSE.
