import { TaskService } from "@/src/gen/proto/backend/v1/task_service-TaskService_connectquery";
import type { RequestHandler } from "msw";
import { connectMock } from "../connectRpc";
import { dummySampleCases } from "../fixtures/sampleCases";

export const taskHandlers: RequestHandler[] = [
  connectMock(TaskService, "getTestcaseSets", (ctx, res, _, encodeResp) => {
    const testcases = [...dummySampleCases];
    return res(
      ctx.delay(500),
      encodeResp({ testcaseSets: [], testcases }),
    );
  }),
];
