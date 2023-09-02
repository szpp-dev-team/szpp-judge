import { createGrpcWebTransport } from "@bufbuild/connect-web";

export const backendGrpcBaseUrl = "http://localhost:8080";
export const backendGrpcTransport = createGrpcWebTransport({ baseUrl: backendGrpcBaseUrl });
