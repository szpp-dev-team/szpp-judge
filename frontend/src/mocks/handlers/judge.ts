import { LangID } from "@/src/gen/langs";
import type { JudgeProgress, SubmissionDetail } from "@/src/gen/proto/backend/v1/judge_resources_pb";
import { JudgeService } from "@/src/gen/proto/backend/v1/judge_service-JudgeService_connectquery";
import { JudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { PlainMessage, Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

const sourcecodeAc = String.raw`#include <iostream>
using namespace std;

int main() {
  int i;
  cin >> i;

  if (i % 15 == 0)
    cout << "Fizz Buzz" << endl;
  else if (i % 3 == 0)
    cout << "Fizz" << endl;
  else if (i % 5 == 0)
    cout << "Buzz" << endl;
  else
    cout << i << endl;

  return 0;
}
`;

const dummySubmissionDetails: PlainMessage<SubmissionDetail>[] = [
  {
    id: 1,
    userId: 1,
    username: "user",
    contestId: 103,
    taskId: 1,
    taskTitle: "すずっぴー君のおつかい",
    score: 100,
    langId: "cpp/20/gcc" satisfies LangID,
    sourceCode: sourcecodeAc,
    status: JudgeStatus.AC,
    execTimeMs: 1,
    execMemoryKib: 3000,
    testcaseResults: [
      { testcaseName: "sample_1", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
      { testcaseName: "sample_2", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
      { testcaseName: "testcase_1", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
      { testcaseName: "testcase_2", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
      { testcaseName: "testcase_3", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
      { testcaseName: "testcase_4", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
    ],
    submittedAt: Timestamp.fromDate(new Date("2023-09-25")),
  },
];

const dummyJudgeProgress: Record<"inProgress" | "ac" | "wa" | "ce", PlainMessage<JudgeProgress>> = {
  inProgress: { status: JudgeStatus.JUDGE_STATUS_UNSPECIFIED, totalTestcases: 20, completedTestcases: 8 },
  ac: { status: JudgeStatus.AC, totalTestcases: 20, completedTestcases: 20 },
  wa: { status: JudgeStatus.WA, totalTestcases: 20, completedTestcases: 12 },
  ce: { status: JudgeStatus.CE, totalTestcases: 20, completedTestcases: 0 },
};

export const judgeHandlers: RequestHandler[] = [
  grpcMock(JudgeService, "getSubmissionDetail", async (ctx, res, decodeReq, encodeResp) => {
    const { id } = await decodeReq();

    const submissionDetail = dummySubmissionDetails.filter(submission => submission.id === id);
    if (submissionDetail.length > 0) {
      return res(encodeResp({ submissionDetail: submissionDetail[0] }));
    } else {
      return res(ctx.status(404));
    }
  }),
  grpcMock(JudgeService, "getJudgeProgress", async (ctx, res, decodeReq, encodeResp) => {
    const { submissionId } = await decodeReq();

    // モックでは submissions は 100 件しかないとする
    if (submissionId > 100) {
      return res(ctx.status(404));
    }

    const r = Math.random();
    const judgeProgress = r < 0.25
      ? dummyJudgeProgress.inProgress
      : r < 0.5
      ? dummyJudgeProgress.ac
      : r < 0.75
      ? dummyJudgeProgress.wa
      : dummyJudgeProgress.ce;

    return res(encodeResp({ judgeProgress }));
  }),
];
