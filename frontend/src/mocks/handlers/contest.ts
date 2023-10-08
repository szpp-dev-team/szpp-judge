import {
  type Contest,
  type ContestTask,
  ContestType,
  StandingsRecord,
  StandingsRecord_TaskDetail,
} from "@/src/gen/proto/backend/v1/contest_resources_pb";
import { ContestService } from "@/src/gen/proto/backend/v1/contest_service-ContestService_connectquery";
import { Duration } from "@/src/util/time";
import { type PlainMessage, Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { connectMock } from "../connectRpc";
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

// StandingsRecord_TaskDetail.untilAc を作るため
const duration = (/** ミリ秒 */ dur: number): PlainMessage<import("@bufbuild/protobuf").Duration> => ({
  seconds: BigInt(Math.trunc(dur / Duration.SECOND)),
  nanos: 0,
});

const groupBy = <K extends PropertyKey, T>(
  objectArray: readonly T[],
  propAsKey: (t: T) => K,
): [Partial<Record<K, T[]>>, ReadonlySet<K>] => {
  const groupKeys = new Set<K>();
  const resultObj = objectArray.reduce<Partial<Record<K, T[]>>>((groups, item) => {
    const key = propAsKey(item); // T[K] extends PropertyKey を保証できないので propAsKey に任せる
    groupKeys.add(key);
    const curGroup = groups[key] ?? [];
    return { ...groups, [key]: [...curGroup, item] };
  }, {});

  return [resultObj, groupKeys];
};

const isTasksViewable = (contest: PlainMessage<Contest>, now: Date): boolean => {
  const startAt = contest.startAt! as Timestamp;
  return now >= startAt.toDate();
};

const generateContest = (slug: string): PlainMessage<Contest> | null => {
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
    default:
      return null;
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

/**
 * 重み付き抽選器
 * @param table key=抽選対象のプロパティ(string),val=割合(number)
 * @returns 選ばれたプロパティ
 * @example
 * ```js
 * weightedPick({ a: 10, b: 20, c: 20 }) // -> a: 20%, b: 40%, c: 40%
 * ```
 */
const weightedPick = <T extends Record<string, number>>(table: T): keyof T => {
  const totalWeight = Object.values(table).reduce((accum, val) => accum + val, 0);
  const sorted = Object.entries(table)
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    .sort(([_k1, a], [_k2, b]) => a - b)
    .reduce<Readonly<Record<string, number>>>((accum, [key, val]) => ({ ...accum, [key]: val }), {});

  const r = 1 + Math.floor(Math.random() * totalWeight);
  let cursor = 0;

  for (const k in sorted) {
    cursor += sorted[k];
    if (r <= cursor) {
      return k;
    }
  }
  return "";
};

const generateStandings = (participantsSize: number, taskSize: number) => {
  taskSize = Math.min(taskSize, contestTasks.length); // contestTasks.length 以上のタスクは作れない

  const AC_SUBMIT_ID = 1; // モックなのでどの提出にも同じ ID を振る

  const standingsNoRank: Omit<PlainMessage<StandingsRecord>, "rank">[] = [];

  for (let user = 0; user < participantsSize; user++) {
    const taskDetailList: PlainMessage<StandingsRecord_TaskDetail>[] = [];

    // 参加者ごとに提出を作る
    for (let j = 0; j < taskSize; j++) {
      const t = contestTasks[j];

      if (t.difficulty > 4 && Math.random() < 0.05) { // true ならこのタスクは未提出
        taskDetailList.push({
          penaltyCount: 0,
          score: 0,
          taskId: t.id,
          acSubmitId: undefined,
          untilAc: undefined,
        });
        continue;
      }

      const [submissionCount, acCount] = (function calcOnDifficulty(diff: typeof t.difficulty): [number, number] {
        const hasAc = weightedPick(diff <= 3 ? { ac: 6, wa: 4 } : { ac: 2, wa: 8 }) === "ac";
        if (hasAc) {
          const key = weightedPick(
            diff <= 3
              ? { oneSubmitOneAc: 6, manySubmitManyAc: 1, manySubmitOneAc: 3 }
              : { oneSubmitOneAc: 8, manySubmitManyAc: 1, manySubmitOneAC: 1 },
          );
          const submissionCount = key === "oneSubmitOneAc" ? 1 : 1 + Math.floor(Math.random() * 3);
          const acCount = key === "manySubmitManyAc" ? 1 + Math.floor(Math.random() * submissionCount) : 1;
          return [submissionCount, acCount];
        } else {
          const submissionCount = 1 + Math.floor(Math.random() * 3);
          return [submissionCount, 0];
        }
      })(t.difficulty);

      const penaltyCount = submissionCount - acCount;

      taskDetailList.push({
        penaltyCount,
        score: acCount > 0 ? t.score : 0, // モックなので部分点はなし. 0, 100
        taskId: t.id,
        acSubmitId: acCount > 0 ? AC_SUBMIT_ID : undefined,
        untilAc: acCount > 0
          ? duration(Math.random() * Duration.HOUR * 2)
          : undefined,
      });
    }

    standingsNoRank.push({
      username: `user${user}`,
      totalScore: taskDetailList.reduce((accum, item) => accum + item.score, 0),
      totalPenaltyCount: taskDetailList.reduce((accum, item) => accum + item.penaltyCount, 0),
      latestAcAt: undefined, // TODO: 削除?
      taskDetailList,
    });
  }

  // 満点を超えていないことを確かめる
  const perfect = contestTasks.slice(0, taskSize).reduce((accum, t) => accum + t.score, 0);
  for (const row of standingsNoRank) {
    if (row.totalScore > perfect) {
      throw new Error(
        "one record of standings exceeded perfect(=max) score ! (perfect="
          + perfect + "but acctual=" + row.totalScore + ")",
      );
    }
  }

  const [rowsGroupedByScore, groupKeys] = groupBy(standingsNoRank, s => s.totalScore); // 得点でグループ化
  const standings = Array.from(groupKeys.values()).sort((a, b) => b - a) // 得点の降順に処理する
    .flatMap(key => rowsGroupedByScore[key]!.sort((a, b) => a.totalPenaltyCount - b.totalPenaltyCount)) // 同点はペナ少がランク高
    .map<PlainMessage<StandingsRecord>>((record, i) => ({
      rank: i + 1,
      ...record,
    }));

  // flat() できていることを確かめるために standings.length が参加人数と同じかどうかを確認する
  if (standings.length !== participantsSize) {
    throw new Error("standings.length does not match with participantsSize");
  }

  return standings;
};

const standingsMap: Map<string, PlainMessage<StandingsRecord>[]> = new Map([
  ["1", generateStandings(0, 9)], // 開始2日前ver
  ["2", generateStandings(0, 9)], // 開始5秒前ver
  ["3", generateStandings(5, 9)], // 開催中ver
  ["4", generateStandings(20, 9)], // 終了5秒前ver
  ["5", generateStandings(50, 9)], // 終了後ver
]);

export const contestHandlers: RequestHandler[] = [
  connectMock(ContestService, "listContests", async (ctx, res, _, encodeResp) => {
    return res(
      ctx.delay(500),
      encodeResp({ contests }),
    );
  }),
  connectMock(ContestService, "getContest", async (ctx, res, decodeReq, encodeResp) => {
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
  connectMock(ContestService, "listContestTasks", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug } = await decodeReq();
    const contest = generateContest(contestSlug);
    if (contest == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    if (!isTasksViewable(contest, new Date())) {
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
  connectMock(ContestService, "getContestTask", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug, taskId } = await decodeReq();
    const contest = generateContest(contestSlug);
    const task = dummyTasks.find((t) => t.id === taskId);

    if (contest == null || task == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    if (!isTasksViewable(contest, new Date())) {
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
  connectMock(ContestService, "getMySubmissionStatuses", async (ctx, res, decodeReq, encodeResp) => {
    const { contestSlug } = await decodeReq();
    const contest = generateContest(contestSlug);
    if (contest == null) {
      return res(
        ctx.delay(500),
        ctx.status(404),
      );
    }
    if (!isTasksViewable(contest, new Date())) {
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
  grpcMock(ContestService, "getStandings", async (ctx, res, decodeReq, encodeResp) => {
    const { contestId } = await decodeReq(); // TODO: proto 更新して contest_slug にしたい
    const standingsList = standingsMap.get(String(contestId));

    if (!standingsList) {
      return res(ctx.status(404));
    } else {
      return res(
        ctx.delay(500),
        encodeResp({ standingsList }),
      );
    }
  }),
];
