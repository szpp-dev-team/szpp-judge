import { listContestTasks } from "@/src/gen/proto/backend/v1/services-ContestService_connectquery";
import { PartialMessage } from "@bufbuild/protobuf";
import { useQuery } from "@tanstack/react-query";
import { ListContestTasksRequest } from "../gen/proto/backend/v1/messages_pb";

export const useListContestTasks = (input: PartialMessage<ListContestTasksRequest>) => {
  const contestTasksQuery = useQuery(listContestTasks.useQuery(input));
  return { contestTasksQuery };
};
