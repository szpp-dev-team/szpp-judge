import {
  calcNthTaskSeq,
  useGetContestTaskResolvingTaskSeq,
  useRouterContestSlug,
  useRouterContestTaskSeq,
} from "@/src/usecases/contest";
import { Box, Card, Heading, Text } from "@chakra-ui/react";
import { TaskDetailCommon } from "../../model/task/TaskDetailCommon";

export const ContestTaskDetail = () => {
  const contestSlug = useRouterContestSlug();
  const taskSeq = useRouterContestTaskSeq();
  const { contest, tasks, task, isContestStarted } = useGetContestTaskResolvingTaskSeq({ contestSlug, taskSeq });
  const seqCode = calcNthTaskSeq(taskSeq - 1, tasks?.length ?? 0);

  if (contest != null && !isContestStarted) {
    return <Text>コンテスト開始前なので問題を表示できません</Text>;
  }

  if (contest == null || tasks == null || task == null) {
    return <Text>読み込み中...</Text>;
  }

  return (
    <Box px={16} h="100%">
      <TaskDetailCommon task={task} seqCode={seqCode} />
    </Box>
  );
};
