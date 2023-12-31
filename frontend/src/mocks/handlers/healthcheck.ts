import { HealthcheckService } from "@/src/gen/proto/backend/v1/healthcheck_service-HealthcheckService_connectquery";
import type { RequestHandler } from "msw";
import { connectMock } from "../connectRpc";

export const healthcheckHandlers: RequestHandler[] = [
  connectMock(HealthcheckService, "ping", async (ctx, res, decodeReq, encodeResp) => {
    const { name } = await decodeReq();
    return res(
      ctx.delay(300),
      encodeResp({ message: `[mock] Hi! ${name} :)` }),
    );
  }),
];
