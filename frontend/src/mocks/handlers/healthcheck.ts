import { HealthcheckService } from "@/src/gen/proto/backend/v1/healthcheck_service-HealthcheckService_connectquery";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

export const healthcheckHandlers: RequestHandler[] = [
  grpcMock(HealthcheckService, "ping", async (ctx, res, decodeReq, encodeResp) => {
    const { name } = await decodeReq();
    return res(
      ctx.delay(300),
      encodeResp({ message: `[mock] Hi! ${name} :)` }),
    );
  }),
];
