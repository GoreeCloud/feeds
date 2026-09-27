# GoreeCloud Feeds Protocol Security

## Current status

This repository contains no released protocol, generated package, or runtime implementation.

## Security requirements for future contracts

Future protocol definitions should:

- avoid embedding reusable credentials, private keys, or protected secrets;
- minimize exposed account and content metadata;
- make authentication and authorization context explicit without allowing clients to self-assert authority;
- define safe error behavior without leaking private server internals;
- validate untrusted identifiers, URLs, content references, and user-controlled fields;
- document replay, ordering, integrity, pagination, and resource-exhaustion considerations where applicable;
- version security-relevant behavior explicitly; and
- coordinate accepted security requirements with GoreeCloud Identity and Wardveil Security.

## Reporting security issues

Do not publish confidential vulnerability details or active secrets in public issues. Use an approved private GitHub security-reporting mechanism or another owner-approved private channel when confidential handling is required.
