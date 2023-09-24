import { type Contest, type ContestTask, ContestType } from "@/src/gen/proto/backend/v1/contest_resources_pb";
import { ContestService } from "@/src/gen/proto/backend/v1/contest_service-ContestService_connectquery";
import { Duration } from "@/src/util/time";
import { type PlainMessage, Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";
import { dummyTasks } from "./task";

const contestTasks: PlainMessage<ContestTask>[] = [
  { ...dummyTasks[0], score: 100 },
  { ...dummyTasks[1], score: 200 },
  { ...dummyTasks[2], score: 314 },
  { ...dummyTasks[3], score: 400 },
  { ...dummyTasks[4], score: 450 },
  { ...dummyTasks[5], score: 500 },
  { ...dummyTasks[6], score: 600 },
  { ...dummyTasks[7], score: 700 },
  { ...dummyTasks[8], score: 1333 },
];

const contestTaskIds = contestTasks.map((task) => task.id);

const contestDescription = String.raw`
## SZPP Beginners Contest とは？
説明説明説明説明説明説明説明説明

## Writer
Bob, Alice, Benjamin
`;

// レスポンスを生成する時点で new Date() をすると、APIを叩くたびに
// コンテスト開始時刻や終了時刻が更新されてしまう。
// すると開始5秒前や終了5秒前の動作確認が難しくなりそう (例えば開始時点で問題一覧取得を叩いたら 403 にならないで欲しい)。
// そのため、このファイルが読み込まれた時点で now をフリーズしている。
// now を更新したい場合はブラウザでページをリロードすればよい。
const now = new Date();
const beforeNow = (dur: number) => Timestamp.fromDate(new Date(now.getTime() - dur));
const afterNow = (dur: number) => Timestamp.fromDate(new Date(now.getTime() + dur));

const canTasksViewable = (contest: PlainMessage<Contest>, now: Date): boolean => {
  const startAt = contest.startAt! as Timestamp;
  return now >= startAt.toDate();
};

const generateContest = (slug: string): PlainMessage<Contest> | undefined => {
  switch (slug) {
    case "sbc001":
      return {
        id: 101,
        slug: "sbc001",
        name: "開始2日前verコンテスト",
        contestType: ContestType.OFFICIAL,
        startAt: afterNow(Duration.DAY * 2),
        endAt: afterNow(Duration.DAY * 2 + Duration.MINUTE * 90),
        penaltySeconds: 300,
        taskIds: contestTaskIds,
        description: contestDescription,
      };
    case "sbc002":
      return {
        id: 102,
        slug: "sbc002",
        name: "開始5秒前verコンテスト",
        contestType: ContestType.OFFICIAL,
        startAt: afterNow(Duration.SECOND * 5),
        endAt: afterNow(Duration.MINUTE * 90),
        penaltySeconds: 300,
        taskIds: contestTaskIds,
        description: contestDescription,
      };
    case "sbc003":
      return {
        id: 103,
        slug: "sbc003",
        name: "開催中verコンテスト",
        contestType: ContestType.OFFICIAL,
        startAt: Timestamp.fromDate(new Date()),
        endAt: afterNow(Duration.MINUTE * 90),
        penaltySeconds: 300,
        taskIds: contestTaskIds,
        description: contestDescription,
      };
    case "sbc004":
      return {
        id: 104,
        slug: "sbc004",
        name: "終了5秒前verコンテスト",
        contestType: ContestType.OFFICIAL,
        startAt: beforeNow(Duration.MINUTE * 90),
        endAt: afterNow(Duration.SECOND * 5),
        penaltySeconds: 300,
        taskIds: contestTaskIds,
        description: contestDescription,
      };
    case "sbc005":
      return {
        id: 105,
        slug: "sbc005",
        name: "終了後verコンテスト/長いタイトル長いタイトル長いタイトル長いタイトル",
        contestType: ContestType.OFFICIAL,
        startAt: beforeNow(Duration.DAY * 2),
        endAt: beforeNow(Duration.DAY * 2 - Duration.MINUTE * 90),
        penaltySeconds: 300,
        taskIds: contestTaskIds,
        description: contestDescription,
      };
  }
};

// TODO: {参加未登録、参加登録済み} の全パターンを追加する
const contests: PlainMessage<Contest>[] = [
  generateContest("sbc001")!,
  generateContest("sbc002")!,
  generateContest("sbc003")!,
  generateContest("sbc004")!,
  generateContest("sbc005")!,
];

export const contestHandlers: RequestHandler[] = [
  grpcMock(ContestService, "listContests", async (ctx, res, _, encodeResp) => {
    return res(
      ctx.delay(500),
      encodeResp({ contests }),
    );
  }),
  grpcMock(ContestService, "getContest", async (ctx, res, decodeReq, encodeResp) => {
    const { slug } = await decodeReq();
    const contest = generateContest(slug);
    if (contest == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    return res(
      ctx.delay(500),
      encodeResp({ contest }),
    );
  }),
  grpcMock(ContestService, "listContestTasks", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug } = await decodeReq();
    const contest = generateContest(contestSlug);
    if (contest == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    if (!canTasksViewable(contest, new Date())) {
      return res(
        ctx.delay(500),
        ctx.status(403),
      );
    }
    return res(
      ctx.delay(500),
      encodeResp({ tasks: contestTasks }),
    );
  }),
  grpcMock(ContestService, "getContestTask", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug, taskId } = await decodeReq();
    const contest = generateContest(contestSlug);
    const task = dummyTasks.find((t) => t.id === taskId);

    if (contest == null || task == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    if (!canTasksViewable(contest, new Date())) {
      return res(
        ctx.delay(500),
        ctx.status(403),
      );
    }
    return res(
      ctx.delay(500),
      encodeResp({ task }),
    );
  }),
  grpcMock(ContestService, "getMySubmissionStatuses", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug } = await decodeReq();
    const contest = generateContest(contestSlug);
    if (contest == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    if (!canTasksViewable(contest, new Date())) {
      return res(
        ctx.delay(500),
        ctx.status(403),
      );
    }
    // 25%: 未提出
    // 25%: 0点
    // 25%: 50点
    // 25%: 満点
    const submissionStatuses = contestTasks.map(t => {
      const r = Math.random();
      return {
        taskId: t.id,
        score: r < 0.25 ? undefined : r < 0.5 ? 0 : r < 0.75 ? 50 : t.score,
      };
    });
    // 最初の4問は全パターン網羅
    submissionStatuses[0].score = undefined;
    submissionStatuses[1].score = 0;
    submissionStatuses[2].score = 50;
    submissionStatuses[3].score = contestTasks[3].score;
    return res(
      ctx.delay(500),
      encodeResp({ submissionStatuses }),
    );
  }),
];
