import {
  useGetContest,
  useGetContestTask,
  useIsContestStarted,
  useListContestTasks,
  useRouterContestSlug,
  useRouterContestTaskSeq,
  useTaskSeqCode,
} from "@/src/usecases/contest";
import { Container, Text } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { TaskDetailCommon } from "../../model/task/TaskDetailCommon";

export const ContestTaskDetail = () => {
  const contestSlug = useRouterContestSlug();
  const taskSeq = useRouterContestTaskSeq();

  const { contest } = useGetContest({ slug: contestSlug });
  const isContestStarted = useIsContestStarted({ contestSlug, now: new Date() });
  const { tasks } = useListContestTasks({ contestSlug }, { isContestStarted });
  const taskId = tasks?.[taskSeq - 1]?.id;
  const { task, sampleCases } = useGetContestTask({ contestSlug, taskId }, { isContestStarted });

  const seqCode = useTaskSeqCode(taskSeq, tasks?.length);

  const router = useRouter();
  const onSubmitSuccess = () => {
    router.push(`/contests/${contestSlug}/submissions?me`);
  };

  if (contest != null && !isContestStarted) {
    return <Text>コンテスト開始前なので問題を表示できません</Text>;
  }

  if (tasks == null || task == null || sampleCases == null) {
    return <Text>読み込み中...</Text>;
  }

  return (
    <Container px={12} maxW="900px" h="100%" centerContent>
      <TaskDetailCommon task={task} sampleCases={sampleCases} taskSeqCode={seqCode} onSubmitSuccess={onSubmitSuccess} />
    </Container>
  );
};
