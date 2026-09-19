# GoreeCloud Feeds Development

## Current development boundary

The project is currently in repository and documentation foundation work. The five initial repositories exist, but no implementation stack or deployable application foundation has been verified yet.

Do not describe planned roadmap capabilities as implemented merely because they are documented.

## Repository ownership

Use the repository whose responsibility owns the change:

- `feeds` — project-wide architecture, roadmap, compatibility, deployment, release, contribution, and coordination documentation.
- `feeds-server` — server implementation.
- `feeds-web` — web client implementation.
- `feeds-protocol` — shared client/server contracts.
- `feeds-shared` — reusable implementation that is genuinely shared across repositories.

Avoid duplicating implementation across repositories when a clear owner exists.

## Branch and pull-request workflow

Material changes should use short-lived topic branches created from the current intended base.

Preferred branch shape:

`<type>/<descriptive-kebab-case-subject>`

Examples include `feature/`, `fix/`, `docs/`, `test/`, `ci/`, `build/`, `chore/`, and `security/`.

Use pull requests for material promotion to the authoritative default branch. Validate the exact candidate head, and re-run affected validation if the candidate changes after review.

## Status integrity

A branch, pull request, successful tool call, merge, package, or repository file does not by itself prove release, deployment, production acceptance, or Stable status.

Documentation must distinguish:

- planned;
- implemented;
- verified;
- released;
- deployed; and
- production-accepted

states when those distinctions are material.

## Toolchain selection

No language, framework, database, package manager, build system, or runtime is selected by this repository foundation.

When implementation begins, select each toolchain at the repository that owns the affected implementation and document the decision before treating it as a project-wide dependency.

## Secrets and local state

Do not commit reusable credentials, private keys, active tokens, production configuration, personal data, or other protected values.

Repository-specific ignore and environment rules should be added when each implementation toolchain is selected.
