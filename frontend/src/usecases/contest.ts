import { listContestTasks } from "@/src/gen/proto/backend/v1/contest_service-ContestService_connectquery";
import { PartialMessage } from "@bufbuild/protobuf";
import { useQuery } from "@tanstack/react-query";
import { ListContestTasksRequest } from "../gen/proto/backend/v1/contest_service_pb";

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

export const useListContestTasks = (input: PartialMessage<ListContestTasksRequest>) => {
  const contestTasksQuery = useQuery(listContestTasks.useQuery(input));
  return { contestTasksQuery };
};
