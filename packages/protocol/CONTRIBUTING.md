# Contributing to GoreeCloud Feeds Protocol

## Scope discipline

Changes in this repository should be limited to contracts that are genuinely shared across GoreeCloud Feeds components.

Do not move server-private persistence models, web-only UI state, or one-off helper types here merely to centralize files.

## Change workflow

1. Start from the current authoritative main branch.
2. Use a short-lived topic branch.
3. State which components consume or provide the proposed contract.
4. Identify whether the change is additive, behavior-changing, or breaking.
5. Document migration/deprecation implications.
6. Update compatibility/versioning documentation when applicable.
7. Validate the exact candidate revision.
8. Use a pull request for material integration.

Until a protocol implementation and release process exists, documentation changes must not be described as a released interface.
