import { createGrpcWebTransport } from "@bufbuild/connect-web";

export const backendGrpcBaseUrl = process.env.NEXT_PUBLIC_ENVOY_BASE_URL!;
export const backendGrpcTransport = createGrpcWebTransport({ baseUrl: backendGrpcBaseUrl });