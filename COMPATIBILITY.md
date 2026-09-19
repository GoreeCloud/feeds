# GoreeCloud Feeds Compatibility

## Current compatibility state

No stable GoreeCloud Feeds protocol version, client compatibility matrix, supported server version, or backward-compatibility guarantee has been established.

## Planned compatibility model

The shared protocol is intended to live in `feeds-protocol` and define:

- API contracts;
- shared data structures;
- synchronization structures;
- error definitions;
- event definitions;
- protocol versioning; and
- client/server compatibility rules.

All current and future clients are intended to communicate with GoreeCloud Feeds Server through the same contract.

## Change rule

A protocol-affecting change should be defined and versioned in `feeds-protocol` before a server or client claims compatibility with it.

Breaking changes should be explicit. Compatibility claims must be tied to verified implementation rather than inferred from repository names or roadmap intent.

## Current matrix

| Component | Verified compatibility status |
| --- | --- |
| `feeds-server` | No stable protocol implementation verified |
| `feeds-web` | No stable protocol implementation verified |
| Future desktop/mobile clients | Not yet created for implementation |
