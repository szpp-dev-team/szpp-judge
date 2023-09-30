import { createGrpcWebTransport } from "@bufbuild/connect-web";

export const backendGrpcBaseUrl = process.env.NEXT_PUBLIC_ENVOY_BASE_URL!;

export const backendGrpcTransportWithOnlyBaseUrl = createGrpcWebTransport({
  baseUrl: backendGrpcBaseUrl,
});
