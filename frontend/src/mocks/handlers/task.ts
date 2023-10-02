import { TaskService } from "@/src/gen/proto/backend/v1/task_service-TaskService_connectquery";
import type { RequestHandler } from "msw";
import { dummySampleCases } from "../fixtures/sampleCases";
import { grpcMock } from "../grpc";

export const taskHandlers: RequestHandler[] = [
  grpcMock(TaskService, "getTestcaseSets", (ctx, res, _, encodeResp) => {
    const testcases = [...dummySampleCases];
    return res(
      ctx.delay(500),
      encodeResp({ testcaseSets: [], testcases }),
    );
  }),
];
