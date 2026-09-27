# GoreeCloud Feeds Protocol Versioning

## Current state

The repository now contains Development contract version `0.1.0-dev`. It is not a released Stable protocol and does not create a supported production compatibility guarantee.

Repository commits remain source-control identities rather than protocol releases. The `0.1.0-dev` identifier exists to keep the first executable server/client contract explicit while Development work proceeds.

## Planned principles

When a protocol versioning model is selected, it should:

- identify breaking versus compatible changes;
- make server/client compatibility understandable;
- permit additive evolution where practical;
- define migration and deprecation behavior;
- avoid silent breaking changes;
- identify exact released protocol artifacts or schemas;
- distinguish protocol version from application release versions; and
- provide capability negotiation where version numbers alone are insufficient.

## Change discipline

A protocol-affecting implementation should be defined here before dependent repositories claim support for it.

A changed branch head, merged pull request, or generated client does not establish a released protocol version without the separate release and compatibility evidence required by GoreeCloud governance.
