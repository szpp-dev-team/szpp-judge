import { backendBaseUrl } from "@/src/config/connectRpc";
import { Message, type MethodInfo, type PlainMessage } from "@bufbuild/protobuf";
import * as msw from "msw";

// gen/**/services_connect.ts 内の各サービスを表現し、同時に methods の型をキャプチャするための型
export type ProtobufService<T extends Readonly<Record<string, Readonly<MethodInfo>>>> = Readonly<{
  typeName: string;
  methods: T;
}>;

// RequestBodyType_3 等の `_3` がついているものは node_modules/.pnpm/msw@1.2.5_typescript@5.1.6/node_modules/msw/lib/index.d.ts
// の rest.post() の型パラメータをコピペしたことを表わす
export type ResponseResolver<
  ReqMessage,
  RespMessage,
  RequestBodyType_3 extends msw.DefaultBodyType = msw.DefaultBodyType,
  Params_3 extends msw.PathParams<keyof Params_3> = msw.PathParams,
  ResponseBody_3 extends msw.DefaultBodyType = msw.DefaultBodyType,
> = (
  ctx: msw.RestContext,
  res: msw.ResponseComposition<ResponseBody_3>,
  decodeReq: () => Promise<ReqMessage>,
  encodeResp: (m: RespMessage) => msw.ResponseTransformer,
  req: msw.RestRequest<RequestBodyType_3, Params_3>,
) => msw.AsyncResponseResolverReturnType<msw.MockedResponse<ResponseBody_3>>;

// 使用例は mocks/handlers/auth.ts などを参照
export const connectMock = <
  T extends Readonly<Record<string, Readonly<MethodInfo>>>,
  K extends keyof T,
  // @ts-expect-error InstanceType<...> は Message<...> に代入できないと怒られる
  TReq extends Message<TReq> = InstanceType<T[K]["I"]>,
  // @ts-expect-error InstanceType<...> は Message<...> に代入できないと怒られる
  TResp extends Message<TResp> = InstanceType<T[K]["O"]>,
>(
  service: ProtobufService<T>,
  methodName: K,
  resolver: ResponseResolver<TReq, PlainMessage<TResp>>,
) => {
  const method = service.methods[methodName];
  return msw.rest.post(`${backendBaseUrl}/${service.typeName}/${method.name}`, (req, res, ctx) => {
    const decodeReq = () => {
      return req.json<TReq>();
    };
    const encodeResp = (m: PlainMessage<TResp>) => {
      return ctx.json(m);
    };
    return resolver(ctx, res, decodeReq, encodeResp, req);
  });
};
