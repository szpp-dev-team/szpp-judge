import type { SubmissionDetail } from "@/src/gen/proto/backend/v1/judge_resources_pb";
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
    langId: "C++",
    sourceCode: sourcecodeAc,
    status: JudgeStatus.AC,
    execTimeMs: 1,
    execMemoryKib: 3000,
    testcaseResults: [
      { testcaseName: "Sample", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
      { testcaseName: "All", judgeStatus: JudgeStatus.AC, execTimeMs: 1, execMemoryKib: 3000 },
    ],
    submittedAt: Timestamp.fromDate(new Date("2023-09-25")),
  },
];

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
];
