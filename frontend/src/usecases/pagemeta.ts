import {
  calcNthTaskSeq,
  useGetContest,
  useGetContestTaskResolvingTaskSeq,
  useRouterContestSlug,
  useRouterContestTaskSeq,
  useRouterSubmissionId,
} from "./contest";

const TITLE_SUFFIX = "SZPP Judge";

export const usePageTitle = (prefix: string): string => {
  return prefix ? `${prefix} | ${TITLE_SUFFIX}` : TITLE_SUFFIX;
};

export const useContestPageTitle = (prefix: string): string => {
  const slug = useRouterContestSlug();
  const { contest } = useGetContest({ slug });
  const name = contest?.name;
  if (name) {
    return `${prefix} - ${name} | ${TITLE_SUFFIX}`;
  } else {
    return `${prefix} | ${TITLE_SUFFIX}`;
  }
};

export const useContestTaskDetailPageTitle = (): string => {
  const contestSlug = useRouterContestSlug();
  const taskSeq = useRouterContestTaskSeq();

  const { contest, tasks, task } = useGetContestTaskResolvingTaskSeq({ contestSlug, taskSeq });
  const seqCode = calcNthTaskSeq(taskSeq - 1, tasks?.length ?? 0);

  const contestName = contest?.name;
  const taskTitle = task?.title;

  if (contestName == null) {
    return TITLE_SUFFIX;
  }
  if (taskTitle == null) {
    return `${seqCode} - ${contestName} | ${TITLE_SUFFIX}`;
  }
  return `${seqCode} - ${taskTitle} - ${contestName} | ${TITLE_SUFFIX}`;
};

export const useSubmissionDetailPageTitle = (prefix: string = "提出 #"): string => {
  const id = useRouterSubmissionId();
  return useContestPageTitle(`${prefix}${id}`);
};
