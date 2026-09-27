# GoreeCloud Feeds Shared Reuse Guidelines

## Governing principle

Move code into GoreeCloud/feeds-shared only when shared ownership materially improves maintainability and at least two real consumers justify the abstraction.

## Admission criteria

A proposed shared component should normally have:

- at least two identified consumers;
- stable enough semantics to justify shared ownership;
- no hidden dependency on one consumer's private runtime;
- tests that exercise shared behavior independently;
- a clear compatibility and versioning impact; and
- a reason that duplication would be more harmful than the shared dependency.

## Rejected patterns

Do not use this repository for:

- speculative future reuse;
- generic “utils” accumulation;
- protocol definitions that belong in feeds-protocol;
- configuration secrets;
- generated build output;
- server-private database structures; or
- browser/UI implementation details.

## Extraction workflow

When extracting existing code into this repository:

1. identify all real consumers;
2. define the ownership boundary;
3. preserve behavior with tests;
4. establish package/version compatibility;
5. update consuming repositories;
6. validate each consumer against the exact shared revision; and
7. remove superseded duplicated code only after dependent adoption is verified.
