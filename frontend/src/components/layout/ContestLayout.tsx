import { ScoreStatus } from "@/src/model/task";
import {
  useGetContest,
  useGetMySubmissionStatuses,
  useIsContestStarted,
  useListContestTasks,
  useRouterContestSlug,
} from "@/src/usecases/contest";
import { type ReactNode, useMemo } from "react";
import { ContestSidebar, ContestSidebarProps } from "../model/contest/ContestSidebar";
import { GLOBAL_HEADER_H } from "../ui/GlobalHeader";
import { WithHeaderFooter } from "./WithHeaderFooter";

export type ContestLayoutProps = {
  children: ReactNode;
};

export const ContestLayout = ({ children }: ContestLayoutProps) => {
  const contestSlug = useRouterContestSlug();

  const { contest } = useGetContest({ slug: contestSlug });
  const isContestStarted = useIsContestStarted({ contestSlug, now: new Date() });
  const { tasks } = useListContestTasks({ contestSlug }, { isContestStarted });
  const { submissionStatuses } = useGetMySubmissionStatuses({ contestSlug }, { isContestStarted });

  const unifiedTasks = useMemo((): ContestSidebarProps["tasks"] => {
    if (tasks == null || submissionStatuses == null) return [];
    return tasks.map((task) => ({
      id: task.id,
      title: task.title,
      scoreStatus: ScoreStatus.fromScore(task.score, submissionStatuses.find((s) => s.taskId === task.id)?.score),
    }));
  }, [tasks, submissionStatuses]);

  return (
    <WithHeaderFooter
      headerProps={{ contestSlug, contestTitle: contest?.name }}
      leftChildren={
        <ContestSidebar
          top={GLOBAL_HEADER_H}
          startAt={contest?.startAt?.toDate()}
          endAt={contest?.endAt?.toDate()}
          now={new Date()}
          slug={contestSlug}
          tasks={unifiedTasks}
        />
      }
    >
      {children}
    </WithHeaderFooter>
  );
};
