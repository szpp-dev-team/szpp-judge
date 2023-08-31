import { backendGrpcBaseUrl } from "@/src/config/grpc";
import { createAsyncIterable, encodeEnvelope, readAllBytes, UniversalServerResponse } from "@bufbuild/connect/protocol";
import { grpcStatusOk, trailerFlag, trailerSerialize } from "@bufbuild/connect/protocol-grpc-web";
import { AnyMessage, Message as ProtobufMessage, MethodInfo } from "@bufbuild/protobuf";
import * as msw from "msw";

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
const grpcResToMSWRes = (grpcResponse: ReturnType<typeof createGrpcResponse>): msw.ResponseTransformer => {
  return async (res) => {
    const iterable = grpcResponse.body;
    if (iterable) {
      res.body = await readAllBytes(iterable, Number.MAX_SAFE_INTEGER);
    }
    if (grpcResponse.header) {
      grpcResponse.header.forEach((value, key) => res.headers.set(key, value));
    }
    res.status = grpcResponse.status;
    return res;
  };
};

// gen/**/services_connect.ts 内の各サービスを表現し、同時に methods の型をキャプチャするための型
export type GrpcService<T extends Record<string, MethodInfo>> = Readonly<{
  typeName: string;
  methods: T;
}>;

export type ResponseResolver<
  RequestBodyType_3 extends msw.DefaultBodyType = msw.DefaultBodyType,
  Params_3 extends msw.PathParams<keyof Params_3> = msw.PathParams,
  ResponseBody_3 extends msw.DefaultBodyType = msw.DefaultBodyType,
> = (
  ctx: msw.RestContext,
  res: msw.ResponseComposition<ResponseBody_3>,
  decodeReq: () => Promise<unknown>,
  encodeResp: (m: AnyMessage) => msw.ResponseTransformer,
  req: msw.RestRequest<RequestBodyType_3, Params_3>,
) => msw.AsyncResponseResolverReturnType<msw.MockedResponse<ResponseBody_3>>;

// 使用例は mocks/handlers/auth.ts などを参照
export const grpcMock = <
  T extends Record<string, MethodInfo>,
  K extends keyof T,
>(
  service: GrpcService<T>,
  methodName: K,
  resolver: ResponseResolver,
) => {
  const method = service.methods[methodName];
  return msw.rest.post(`${backendGrpcBaseUrl}/${service.typeName}/${method.name}`, (req, res, ctx) => {
    const decodeReq = async () => {
      const a = await req.arrayBuffer();
      return method.I.fromBinary(new Uint8Array(a.slice(5))); // 先頭5バイトのマジックコードを削除
    };
    const encodeResp = (m: AnyMessage) => {
      return grpcResToMSWRes(createGrpcResponse(m));
    };
    return resolver(ctx, res, decodeReq, encodeResp, req);
  });
};
