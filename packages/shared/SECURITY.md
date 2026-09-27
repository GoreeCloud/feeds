# GoreeCloud Feeds Shared Security

## Current status

This repository currently contains no executable library, dependency manifest, package, or runtime code.

## Future shared-code security requirements

Shared implementation should:

- avoid reusable credentials, private keys, tokens, or protected runtime configuration;
- validate untrusted data at appropriate boundaries;
- avoid silently weakening authorization or privacy behavior in consuming applications;
- minimize dependencies and review their security/licensing impact;
- keep generated artifacts traceable to source;
- document security-relevant behavior that consumers must enforce; and
- be validated independently and in each affected consumer when security-sensitive behavior changes.

## Reporting security issues

Do not publish active secrets, private user information, or restricted vulnerability details in public issues. Use an approved private GitHub security-reporting mechanism or another owner-approved private channel when confidential handling is required.
