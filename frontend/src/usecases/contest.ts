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
import { Duration } from "../util/time";

const CHAR_CODE_A = "A".charCodeAt(0);

/**
 * コンテストの問題数に応じて問題番号をアルファベットまたは整数で生成する。
 * @param i 問題の添字 (0-indexed)
 * @param numTotalTasks コンテストの問題数
 */
export const calcNthTaskSeq = (i: number, numTotalTasks: number): string => {
  if (numTotalTasks < 26) {
    return String.fromCharCode(CHAR_CODE_A + i);
  }
  return String(i + 1);
};

export const useRouterContestSlug = () => {
  const { query } = useRouter();
  /// URLに [contest_slug] が含まれていないページでは undefined になるので注意
  return query.contest_slug as string;
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

export const useListContestTasks = (input?: PlainMessage<ListContestTasksRequest>) => {
  const { data, error, isLoading } = useQuery({
    ...listContestTasks.useQuery(input),
    staleTime: 5 * Duration.MINUTE,
  });
  const tasks = data?.tasks;
  return { tasks, error, isLoading };
};

export const useGetContestTask = (input?: PlainMessage<GetContestTaskRequest>) => {
  const { data, error, isLoading } = useQuery({
    ...getContestTask.useQuery(input),
    staleTime: 5 * Duration.MINUTE,
  });
  const task = data?.task;
  return { task, error, isLoading };
};

export const useGetMySubmissionStatuses = (input?: PlainMessage<GetMySubmissionStatusesRequest>) => {
  const { data, error, isLoading } = useQuery({
    ...getMySubmissionStatuses.useQuery(input),
    staleTime: Duration.MINUTE,
  });
  const submissionStatuses = data?.submissionStatuses;
  return { submissionStatuses, error, isLoading };
};
