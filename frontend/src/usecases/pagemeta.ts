import { useGetContest, useRouterContestSlug, useRouterSubmissionId } from "./contest";

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

export const useSubmissionDetailPageTitle = (prefix: string = "提出 #"): string => {
  const id = useRouterSubmissionId();
  return useContestPageTitle(`${prefix}${id}`);
};
