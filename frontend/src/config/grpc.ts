import { createGrpcWebTransport } from "@bufbuild/connect-web";

export const backendGrpcTransport = createGrpcWebTransport({ baseUrl: "http://localhost:8080" });
