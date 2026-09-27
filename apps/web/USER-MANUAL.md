# GoreeCloud Feeds Web User Manual

## Current availability

There is currently no supported GoreeCloud Feeds Web build or deployment for end users.

A Development TypeScript toolchain exists for engineering validation, but it does not render a user interface or connect to GoreeCloud Feeds Server.

## Developer-only validation

With Node.js 24.21.0 LTS and npm 11.19.0:

```sh
npm install --ignore-scripts --no-audit --no-fund --package-lock=false
npm run typecheck
npm run build
```

The generated Development module is not a browser application, production asset bundle, or deployment artifact.

## Unsupported user areas

No application URL, login flow, feed reader, Home/Today interface, server connection, offline cache, installable web-app behavior, browser support matrix, Glaze UI-rendered interface, or production deployment exists yet.

## Planned user areas

When implementation exists, this manual will document only verified behavior for connecting/signing in; subscribing to and navigating feeds; Home and Today views; reading articles; unread/read state; saved articles and favorites; history; search; folders and tags; Smart Feeds and rules; offline behavior; synchronization/connection state; notifications/settings; administration where authorized; and accessibility/keyboard/touch controls.

Do not infer usable UI behavior from roadmap or Development toolchain files alone.
