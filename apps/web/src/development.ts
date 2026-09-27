export const developmentFoundation = Object.freeze({
  product: "GoreeCloud Feeds Web",
  lifecycle: "development",
  protocolTarget: "0.1.0-dev",
  glazeUiTarget: "1.6.0",
  renderedApplicationImplemented: false,
} as const);

export type DevelopmentFoundation = typeof developmentFoundation;
