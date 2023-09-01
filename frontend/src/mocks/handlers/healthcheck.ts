import { PingRequest, PingResponse } from "@/src/gen/proto/backend/v1/messages_pb";
import { HealthcheckService } from "@/src/gen/proto/backend/v1/services_connect";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

export const healthcheckHandlers: RequestHandler[] = [
  grpcMock(HealthcheckService, "ping", async (ctx, res, decodeReq, encodeResp) => {
    const { name } = await decodeReq() as PingRequest;
    return res(
      ctx.delay(300),
      encodeResp(new PingResponse({ message: `[mock] Hi! ${name} :)` })),
    );
  }),
];
