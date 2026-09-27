# Contributing to GoreeCloud Feeds

## Scope

Contributions should preserve the Feeds product-family monorepo boundaries and the distinction between planned and implemented functionality.

## Choose the correct component

- Product-wide documentation and coordination: repository root and `docs/`
- Server implementation: `services/server/`
- Web client: `apps/web/`
- Shared protocol contracts: `packages/protocol/`
- Genuinely reusable internal implementation: `packages/shared/`

A coherent pull request may update multiple component paths when a contract or product behavior requires coordinated change.

## Change workflow

1. Start from the current authoritative target branch.
2. Use a short-lived, purpose-specific topic branch.
3. Keep the change limited to one coherent scope.
4. Update directly affected documentation.
5. Do not include secrets, active credentials, private keys, tokens, or protected operational data.
6. Validate the exact candidate revision.
7. Use a pull request for material integration.
8. Treat merge, release, deployment, and Stable status as separate states.

## Documentation integrity

Roadmap text describes intended capabilities unless implementation evidence establishes otherwise.

Do not convert planned behavior into present-tense implementation claims without verifying the relevant source and tests.

## Coordinated component changes

Protocol changes should be owned in `packages/protocol/`. A single pull request may also update `services/server/` and `apps/web/` against that contract, allowing one atomic review while preserving explicit ownership boundaries.
