import { ContestTask } from "@/src/gen/proto/backend/v1/contest_resources_pb";
import { ContestService } from "@/src/gen/proto/backend/v1/contest_service-ContestService_connectquery";
import { ListContestTasksRequest, ListContestTasksResponse } from "@/src/gen/proto/backend/v1/contest_service_pb";
import { Difficulty } from "@/src/gen/proto/backend/v1/task_resources_pb";
import { PlainMessage } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

const contestTasks: PlainMessage<ContestTask>[] = /* dprint-ignore */ [
  {id:98,title:"すずっぴー君のおつかい",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.BEGINNER,score:100},
  {id:99,title:"すずっぴー君の怒り",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.BEGINNER,score:200},
  {id:100,title:"すずっぴー君の分裂",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.EASY,score:300},
  {id:101,title:"すずっぴー君の爆散",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.EASY,score:314},
  {id:102,title:"師匠との出会い",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.MEDIUM,score:400},
  {id:103,title:"師匠の爆散",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.BEGINNER,score:500},
  {id:104,title:"すずっぴー君の凝固",execTimeLimit:2000,execMemoryLimit:1024,difficulty:Difficulty.HARD,score:600},
  {id:105,title:"長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル",execTimeLimit:2000,execMemoryLimit:0,difficulty:Difficulty.IMPOSSIBLE,score:1333},
];

export const contestHandlers: RequestHandler[] = [
  grpcMock(ContestService, "listContestTasks", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug } = await decodeReq() as ListContestTasksRequest;
    const found = contestSlug === "sbc001";
    if (!found) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    return res(
      ctx.delay(500),
      encodeResp(
        new ListContestTasksResponse({
          tasks: contestTasks,
        }),
      ),
    );
  }),
];
