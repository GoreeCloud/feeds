# Contributing to GoreeCloud Feeds Shared

## Scope discipline

Contributions must preserve the repository's narrow shared-implementation role.

Before adding code, identify the real consumers and explain why the implementation belongs here instead of in an owning application/service repository.

## Change workflow

1. Start from the current authoritative main branch.
2. Use a short-lived topic branch.
3. Identify all affected consumers.
4. Add or update independent shared tests when implementation exists.
5. Document compatibility or migration impact.
6. Validate the exact candidate revision.
7. Validate affected consumers before claiming compatible adoption.
8. Use a pull request for material integration.

Do not introduce speculative abstractions or protocol contracts here.
