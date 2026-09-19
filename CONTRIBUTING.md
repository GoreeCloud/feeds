# Contributing to GoreeCloud Feeds

## Scope

Contributions should preserve the defined repository boundaries and the distinction between planned and implemented functionality.

## Choose the correct repository

- Project-wide documentation and coordination: `feeds`
- Server implementation: `feeds-server`
- Web client: `feeds-web`
- Shared protocol contracts: `feeds-protocol`
- Genuinely reusable internal code: `feeds-shared`

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

Do not convert planned behavior into present-tense implementation claims without verifying the owning repository and the relevant tests or runtime evidence.

## Cross-repository changes

When one change affects multiple repositories, identify the ownership and dependency order explicitly. Protocol changes should be coordinated through `feeds-protocol` before dependent clients or servers claim compatibility with them.
