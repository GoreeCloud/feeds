# GoreeCloud Feeds Deployment

## Current status

There is no verified production GoreeCloud Feeds deployment, supported release package, Release Candidate, or Stable artifact.

Development source now includes a Go server executable foundation under `services/server/`, a TypeScript web build foundation under `apps/web/`, a PostgreSQL 18 persistence target, and monorepo CI. Those facts do not establish a supported deployment.

The following remain unverified or unestablished for production use:

- supported container or package artifact;
- production hostname and network exposure;
- production authentication and authorization;
- complete runtime wiring of retrieval, parsing, persistence, and synchronization;
- production database configuration and upgrade procedure;
- media/search storage topology;
- backup and restore acceptance;
- supported upgrade and rollback procedure; and
- production acceptance evidence.

## Deployment boundary

The server is the intended authoritative service and clients consume its shared protocol.

Any future deployment model must account for server runtime, durable article/account data, media cache, search state, authentication, network exposure, protocol compatibility, backup/recovery, migration, upgrade, and rollback behavior.

## Completion rule

Deployment documentation may publish operational commands, images, ports, environment variables, storage paths, migration procedures, or guarantees only after those values are implemented and verified in this repository.
