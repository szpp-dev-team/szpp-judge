import {
  useGetContest,
  useGetContestTask,
  useIsContestStarted,
  useListContestTasks,
  useRouterContestSlug,
  useRouterContestTaskSeq,
  useTaskSeqCode,
} from "@/src/usecases/contest";
import { Box, Text } from "@chakra-ui/react";
import { TaskDetailCommon } from "../../model/task/TaskDetailCommon";

export const ContestTaskDetail = () => {
  const contestSlug = useRouterContestSlug();
  const taskSeq = useRouterContestTaskSeq();

  const { contest } = useGetContest({ slug: contestSlug });
  const isContestStarted = useIsContestStarted({ contestSlug, now: new Date() });
  const { tasks } = useListContestTasks({ contestSlug }, { isContestStarted });
  const taskId = tasks?.[taskSeq - 1]?.id;
  const { task } = useGetContestTask({ contestSlug, taskId }, { isContestStarted });

  const seqCode = useTaskSeqCode(taskSeq, tasks?.length);

  if (contest != null && !isContestStarted) {
    return <Text>コンテスト開始前なので問題を表示できません</Text>;
  }

  if (contest == null || tasks == null || task == null) {
    return <Text>読み込み中...</Text>;
  }

  return (
    <Box px={12} h="100%">
      <TaskDetailCommon task={task} seqCode={seqCode} />
    </Box>
  );
};
