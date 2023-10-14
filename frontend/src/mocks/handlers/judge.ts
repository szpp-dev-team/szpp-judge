import { LangID } from "@/src/gen/langs";
import type { JudgeProgress, SubmissionDetail, SubmissionSummary } from "@/src/gen/proto/backend/v1/judge_resources_pb";
import { JudgeService } from "@/src/gen/proto/backend/v1/judge_service-JudgeService_connectquery";
import { JudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { PlainMessage, Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { connectMock } from "../connectRpc";

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
  {
    id: 2,
    userId: 1,
    username: "user",
    contestId: 103,
    taskId: 1,
    taskTitle: "すずっぴー君のおつかい",
    score: 0,
    langId: "cpp/20/gcc" satisfies LangID,
    sourceCode: sourcecodeAc,
    status: JudgeStatus.WJ,
    execTimeMs: 1,
    execMemoryKib: 3000,
    testcaseResults: [
      { testcaseName: "sample_1", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
    ],
    submittedAt: Timestamp.fromDate(new Date("2023-09-25")),
  },
];

const dummySubmissionSummaries: PlainMessage<SubmissionSummary>[] = dummySubmissionDetails.map((submissions) => {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const { testcaseResults, sourceCode, status, ...summary } = submissions; // 不要なプロパティを除く
  return { judgeStatus: status, ...summary }; // judgeStatus は SubmissionDetail.status と名前が違うだけで同じなので流用
});

const dummyJudgeProgress: Record<"wj" | "ac" | "wa" | "ce", PlainMessage<JudgeProgress>> = {
  wj: { status: JudgeStatus.WJ, totalTestcases: 20, completedTestcases: 0 },
  ac: { status: JudgeStatus.AC, totalTestcases: 20, completedTestcases: 20 },
  wa: { status: JudgeStatus.WA, totalTestcases: 20, completedTestcases: 12 },
  ce: { status: JudgeStatus.CE, totalTestcases: 20, completedTestcases: 0 },
};

/** 4回目の getJudgeProgress で AC や WA を返すためにリクエスト回数を記録しておく辞書 */
const reqCountMap = new Map<number, number>();
/** 同じsubmissionId のリクエストに対して確定したジャッジ結果を返すために記録しておく辞書 */
const fixedResultMap = new Map<number, PlainMessage<JudgeProgress>>();

export const judgeHandlers: RequestHandler[] = [
  connectMock(JudgeService, "submit", async (ctx, res, _, encodeResp) => {
    return res(ctx.delay(500), encodeResp({ submissionId: 1 }));
  }),
  connectMock(JudgeService, "getSubmissionDetail", async (ctx, res, decodeReq, encodeResp) => {
    const { id } = await decodeReq();

    const submissionDetail = dummySubmissionDetails.filter(submission => submission.id === id);
    if (submissionDetail.length > 0) {
      return res(encodeResp({ submissionDetail: submissionDetail[0] }));
    } else {
      return res(ctx.status(404));
    }
  }),
  connectMock(JudgeService, "listSubmissions", async (ctx, res, decodeReq, encodeResp) => {
    const { username } = await decodeReq(); // デバッグの都合で contestId を無視

    const submissionSummaries = dummySubmissionSummaries.filter(s => {
      if (username && s.username !== username) {
        return false;
      } else {
        return true;
      }
    });

    return res(ctx.delay(500), encodeResp({ submissions: submissionSummaries }));
  }),
  connectMock(JudgeService, "getJudgeProgress", async (ctx, res, decodeReq, encodeResp) => {
    const { submissionId } = await decodeReq();

    // モックでは submissions は 100 件しかないとする
    if (submissionId > 100) {
      return res(ctx.status(404));
    }

    reqCountMap.set(submissionId, (reqCountMap.get(submissionId) ?? 0) + 1);
    const c = reqCountMap.get(submissionId)!;
    if (c < 4) {
      return res(encodeResp({ judgeProgress: dummyJudgeProgress.wj }));
    } else if (c === 4) {
      const r = Math.random();
      const judgeProgress = r < 0.33 ? dummyJudgeProgress.ac : r < 0.66 ? dummyJudgeProgress.ce : dummyJudgeProgress.wa;
      fixedResultMap.set(submissionId, judgeProgress);
      return res(encodeResp({ judgeProgress }));
    } else {
      const r = fixedResultMap.get(submissionId);
      return res(encodeResp({ judgeProgress: r ?? dummyJudgeProgress.wj })); // WJ になるはずはない
    }
  }),
];
