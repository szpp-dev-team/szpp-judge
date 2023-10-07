import {
  getContest,
  getContestTask,
  getMySubmissionStatuses,
  listContests,
  listContestTasks,
} from "@/src/gen/proto/backend/v1/contest_service-ContestService_connectquery";
import type {
  GetContestRequest,
  GetMySubmissionStatusesRequest,
  ListContestsRequest,
  ListContestTasksRequest,
} from "@/src/gen/proto/backend/v1/contest_service_pb";
import { PlainMessage } from "@bufbuild/protobuf";
import { useQuery } from "@tanstack/react-query";
import { useRouter } from "next/router";
import { useMemo } from "react";
import { Duration } from "../util/time";

const CHAR_CODE_A = "A".charCodeAt(0);

/**
 * コンテストの問題数に応じて問題番号コードをアルファベットまたは整数で生成する。
 * @param seq 問題の連番 **(1-indexed)**
 * @param numContestTasks コンテストの問題数
 */
export const useTaskSeqCode = (seq: number, numContestTasks: number | undefined): string => {
  return useMemo(() => {
    if (1 <= seq && seq <= 26 && numContestTasks != null && numContestTasks <= 26) {
      return String.fromCharCode(CHAR_CODE_A + seq - 1); // seq は 1-indexed なので -1 する
    }
    return String(seq);
  }, [seq, numContestTasks]);
};

/**
 * コンテストの問題数に応じて問題番号コードの配列をアルファベットまたは整数で生成する。
 * @param numContestTasks コンテストの問題数
 */
export const useTaskSeqCodes = (numContestTasks: number): string[] => {
  return useMemo(() => {
    if (numContestTasks <= 26) {
      return [...Array(numContestTasks)].map((_, i) => String.fromCharCode(CHAR_CODE_A + i));
    }
    return [...Array(numContestTasks)].map((_, i) => String(i + 1));
  }, [numContestTasks]);
};

/**
 * URL から contest_slug を取得
 * URL に [contest_slug] が含まれていないページでは undefined になるので注意
 */
export const useRouterContestSlug = () => {
  const { query } = useRouter();
  return query.contest_slug as string;
};

/**
 * URL から contest_task_seq を取得
 * URL に [contest_task_seq] が含まれていないページでは undefined になるので注意
 * contest_task_seq が非数値の場合は NaN になりえるので注意
 */
export const useRouterContestTaskSeq = () => {
  const { query } = useRouter();
  return Number.parseInt(query.contest_task_seq as string, 10);
};

/**
 * URL から 提出 ID を取得
 * URL に [submission_id] が含まれていないページでは undefined になるので注意
 */
export const useRouterSubmissionId = () => {
  const { query } = useRouter();
  return query.submission_id as string;
};

const STALE_TIME = {
  // 運営側がコンテストの終了を延長したり問題を追加したりする可能性があるので長すぎない値に
  getContest: Duration.MINUTE,
  listContestTasks: Duration.MINUTE,
  // 自分の提出状況は submit の useMutation によって最新の値にできるが、それはそのWebブラウザのタブ内での話である。
  // 問題ごとにブラウザタブを開いている場合もあるので、例えば問題Aへ提出後に別の問題Bのタブへ切り替えたとき、
  // 提出結果をリフェッチしたいので、短めの値に設定。
  getMySubmissionStatuses: Duration.SECOND * 30,

  // NOTE: 問題文は滅多に変わらないことを仮定している。
  getContestTask: Duration.MINUTE * 30,
} as const;

export const useListContests = (input?: PlainMessage<ListContestsRequest>) => {
  const { data, error, isLoading } = useQuery(listContests.useQuery(input));
  const contests = data?.contests;
  return { contests, error, isLoading };
};

export const useGetContest = (input?: PlainMessage<GetContestRequest>) => {
  const { data, error, isLoading } = useQuery({
    ...getContest.useQuery(input),
    staleTime: STALE_TIME.getContest,
  });
  const contest = data?.contest;
  return { contest, error, isLoading };
};

export const useIsContestStarted = ({ contestSlug, now }: { contestSlug: string; now: Date }): boolean => {
  const { data } = useQuery({
    ...getContest.useQuery({ slug: contestSlug }),
    staleTime: STALE_TIME.getContest,
    select: (({ contest }) => contest != null && contest.startAt!.toDate() < now),
  });
  return data ?? false;
};

export const useListContestTasks = (
  input: PlainMessage<ListContestTasksRequest>,
  opt: { isContestStarted: boolean },
) => {
  const { data, error, isLoading } = useQuery({
    ...listContestTasks.useQuery(input),
    staleTime: STALE_TIME.listContestTasks,
    enabled: opt.isContestStarted,
  });
  const tasks = data?.tasks;
  return { tasks, error, isLoading };
};

export const useGetMySubmissionStatuses = (
  input: PlainMessage<GetMySubmissionStatusesRequest>,
  opt: { isContestStarted: boolean },
) => {
  const { data, error, isLoading } = useQuery({
    ...getMySubmissionStatuses.useQuery(input),
    staleTime: STALE_TIME.getMySubmissionStatuses,
    enabled: opt.isContestStarted,
  });
  const submissionStatuses = data?.submissionStatuses;
  return { submissionStatuses, error, isLoading };
};

/** `taskId != null && isContestStarted` の場合に限り useQuery を有効化する */
export const useGetContestTask = (
  input: { contestSlug: string; taskId: number | undefined },
  opt: { isContestStarted: boolean },
) => {
  const { contestSlug, taskId } = input;

  const { data, error, isLoading } = useQuery({
    ...getContestTask.useQuery({ contestSlug, taskId }),
    staleTime: STALE_TIME.getContestTask,
    enabled: opt.isContestStarted && taskId != null,
  });
  const task = data?.task;
  const sampleCases = data?.samples;
  return { task, sampleCases, error, isLoading };
};
