import {
  useGetContest,
  useGetContestTask,
  useIsContestStarted,
  useListContestTasks,
  useRouterContestSlug,
  useRouterContestTaskSeq,
  useTaskSeqCode,
} from "@/src/usecases/contest";
import { Container } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { TaskDetailCommon } from "../../model/task/TaskDetailCommon";
import { LoadingPane } from "../../ui/LoadingPane";

export const ContestTaskDetail = () => {
  const contestSlug = useRouterContestSlug();
  const taskSeq = useRouterContestTaskSeq();

  const { contest } = useGetContest({ slug: contestSlug });
  const isContestStarted = useIsContestStarted({ contestSlug, now: new Date() });
  const { tasks } = useListContestTasks({ contestSlug }, { isContestStarted });
  const contestTask = tasks?.[taskSeq - 1];
  const taskId = contestTask?.id;
  const { task, sampleCases } = useGetContestTask({ contestSlug, taskId }, { isContestStarted });

  const seqCode = useTaskSeqCode(taskSeq, tasks?.length);

  const router = useRouter();
  const onSubmitSuccess = () => {
    router.push(`/contests/${contestSlug}/submissions?me`);
  };

  if (contest == null || tasks == null || task == null || sampleCases == null) {
    return <LoadingPane />;
  }

  // FIXME: contestTask の execTimeLimit, execMemLimit を使う

  return (
    <Container px={12} maxW="900px" h="100%" centerContent>
      <TaskDetailCommon
        contestId={contest.id}
        task={task}
        score={contestTask!.score}
        sampleCases={sampleCases}
        taskSeqCode={seqCode}
        onSubmitSuccess={onSubmitSuccess}
      />
    </Container>
  );
};
