import {
  getContest,
  getContestTask,
  getMySubmissionStatuses,
  listContests,
  listContestTasks,
} from "@/src/gen/proto/backend/v1/contest_service-ContestService_connectquery";
import type {
  GetContestRequest,
  GetContestTaskRequest,
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

type UseQueryOption = {
  enabled?: boolean;
};

export const useListContests = (input?: PlainMessage<ListContestsRequest>) => {
  const { data, error, isLoading } = useQuery(listContests.useQuery(input));
  const contests = data?.contests;
  return { contests, error, isLoading };
};

export const useGetContest = (input?: PlainMessage<GetContestRequest>) => {
  const { data, error, isLoading } = useQuery({
    ...getContest.useQuery(input),
    // 運営側がコンテストの終了を延長する可能性があるので長過ぎない値に
    staleTime: Duration.MINUTE,
  });
  const contest = data?.contest;
  return { contest, error, isLoading };
};

export const useListContestTasks = (input?: PlainMessage<ListContestTasksRequest>, opt?: UseQueryOption) => {
  const { data, error, isLoading } = useQuery({
    ...listContestTasks.useQuery(input),
    staleTime: 5 * Duration.MINUTE,
    ...opt,
  });
  const tasks = data?.tasks;
  return { tasks, error, isLoading };
};

export const useGetContestTask = (input?: PlainMessage<GetContestTaskRequest>, opt?: UseQueryOption) => {
  const { data, error, isLoading } = useQuery({
    ...getContestTask.useQuery(input),
    staleTime: 5 * Duration.MINUTE,
    ...opt,
  });
  const task = data?.task;
  return { task, error, isLoading };
};

export const useGetContestTaskResolvingTaskSeq = ({ contestSlug, taskSeq }: {
  contestSlug: string;
  taskSeq: number;
}) => {
  const { contest } = useGetContest({ slug: contestSlug });

  const isContestStarted = contest != null && contest.startAt!.toDate()! < new Date();
  const { tasks } = useListContestTasks({ contestSlug }, { enabled: isContestStarted });

  const taskId = tasks?.[taskSeq - 1].id;
  const { task } = useGetContestTask(
    { contestSlug, taskId: taskId! },
    { enabled: taskId != null },
  );

  return { contest, tasks, task, isContestStarted };
};

export const useGetMySubmissionStatuses = (
  input?: PlainMessage<GetMySubmissionStatusesRequest>,
  opt?: UseQueryOption,
) => {
  const { data, error, isLoading } = useQuery({
    ...getMySubmissionStatuses.useQuery(input),
    staleTime: Duration.MINUTE,
    ...opt,
  });
  const submissionStatuses = data?.submissionStatuses;
  return { submissionStatuses, error, isLoading };
};
