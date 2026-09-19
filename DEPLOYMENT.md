# GoreeCloud Feeds Deployment

## Current status

There is currently **no verified deployable GoreeCloud Feeds artifact**.

The repository foundation does not establish:

- a server executable;
- a container image;
- a web-client build;
- a package;
- a deployment manifest;
- a production hostname;
- a listening port;
- a database engine;
- a storage path;
- a supported upgrade path; or
- a production deployment.

## Planned deployment boundary

The roadmap defines GoreeCloud Feeds Server as the authoritative service and clients as consumers of its shared API and synchronization contract.

A future deployment model may therefore need to account for:

- server runtime;
- persistent article and account data;
- media cache;
- backups and restore;
- search storage;
- authentication and authorization;
- network exposure;
- client/server protocol compatibility; and
- upgrade and rollback behavior.

These are planning boundaries, not selected implementation details.

## Completion rule

Deployment documentation should only publish commands, ports, images, environment variables, storage paths, migration procedures, or operational guarantees after those values are established and verified in the owning implementation repository.
