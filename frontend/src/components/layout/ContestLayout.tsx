import { ScoreStatus } from "@/src/model/task";
import {
  useGetContest,
  useGetMySubmissionStatuses,
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
  const slug = useRouterContestSlug();

  const { contest } = useGetContest({ slug });
  const { tasks } = useListContestTasks({ contestSlug: slug });
  const { submissionStatuses } = useGetMySubmissionStatuses({ contestSlug: slug });

  const unifiedTasks = useMemo((): ContestSidebarProps["tasks"] => {
    if (tasks == null || submissionStatuses == null) return [];
    return submissionStatuses.map((s, i) => ({
      id: s.taskId,
      title: tasks[i].title,
      scoreStatus: ScoreStatus.fromScore(tasks[i].score, s.score),
    }));
  }, [tasks, submissionStatuses]);

  return (
    <WithHeaderFooter
      headerProps={{ contestSlug: slug, contestTitle: contest?.name }}
      leftChildren={
        <ContestSidebar
          top={GLOBAL_HEADER_H}
          startAt={contest?.startAt?.toDate()}
          endAt={contest?.endAt?.toDate()}
          now={new Date()}
          slug={slug}
          tasks={unifiedTasks}
        />
      }
    >
      {children}
    </WithHeaderFooter>
  );
};
