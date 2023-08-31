import { AuthService, HealthcheckService } from "@/src/gen/proto/backend/v1/services_connect";
import { createAsyncIterable, encodeEnvelope, readAllBytes, UniversalServerResponse } from "@bufbuild/connect/protocol";
import { grpcStatusOk, trailerFlag, trailerSerialize } from "@bufbuild/connect/protocol-grpc-web";
import { Message as ProtobufMessage, Timestamp } from "@bufbuild/protobuf";
import { RequestHandler as MSWRequestHandler, ResponseTransformer, rest } from "msw";

// Thanks to this blog and this issue.
// Original idea was written for protobuf-ts, and I ported it to `connect` for this file.
// - https://github.com/mswjs/msw/issues/238
// - https://lucas-levin.com/code/blog/mocking-grpc-web-requests-for-integration-testing

/**
 * create a dummy Response from server on transport level. Only for **unaly** call!
 *
 * @see [transport.spec.ts](https://github.com/connectrpc/connect-es/blob/19e78837a094ce6c99af50cef64f418495632c93/packages/connect/src/protocol-grpc-web/transport.spec.ts)
 */
const createGrpcResponse = <Response extends ProtobufMessage>(message: Response): UniversalServerResponse => {
  return {
    status: 200,
    header: new Headers({
      "Content-Type": "application/grpc-web+proto",
    }),
    body: createAsyncIterable([
      encodeEnvelope(0, message.toBinary()),
      encodeEnvelope(
        trailerFlag,
        trailerSerialize(
          new Headers({
            "Grpc-Status": grpcStatusOk,
            "Grpc-Message": "",
          }),
        ),
      ),
    ]),
    trailer: new Headers(),
  };
};

/**
 * Create a transformer that transforms GrpcResponse to MSW compatible response.
 */
const grpcResToMSWRes = (grpcResponse: ReturnType<typeof createGrpcResponse>): ResponseTransformer => {
  return async (res) => {
    const iterable = grpcResponse.body;
    console.log("grpcResponse.body:", iterable);
    if (iterable) {
      res.body = await readAllBytes(iterable, Number.MAX_SAFE_INTEGER);
      console.log("res.body:", res.body);
    }
    if (grpcResponse.header) {
      // eslint-disable-next-line no-unused-vars
      grpcResponse.header.forEach((value, key, _) => res.headers.set(key, value));
    }
    res.status = grpcResponse.status;
    return res;
  };
};

const baseUrl = "http://localhost:8080/";

export const handlers: MSWRequestHandler[] = [
  // eslint-disable-next-line no-unused-vars
  rest.post("http://localhost:8080/backend.v1.HealthcheckService/Ping", (_req, res, _ctx) => {
    const respCls = HealthcheckService.methods.ping.O;
    const message = createGrpcResponse(
      respCls.fromJson({ message: "Hi Bob! Pong :)" }),
    );
    return res(grpcResToMSWRes(message));
  }),

  rest.post(`${baseUrl}${AuthService.typeName}/Login`, async (req, res, ctx) => {
    const ReqClass = AuthService.methods.login.I;
    const RespClass = AuthService.methods.login.O;

    const a = await req.arrayBuffer();
    const { username, password } = ReqClass.fromBinary(new Uint8Array(a.slice(5)));

    const ok = (username === "user" || username === "admin") && password === "pass";
    if (!ok) {
      return res(
        ctx.delay(500),
        ctx.status(401), // 401 Forbidden
      );
    }
    return res(
      ctx.delay(500),
      grpcResToMSWRes(createGrpcResponse(
        new RespClass({
          user: {
            id: 1,
            username,
            isAdmin: username === "admin",
            createdAt: Timestamp.now(),
          },
        }),
      )),
    );
  }),
];
