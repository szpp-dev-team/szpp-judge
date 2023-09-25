import { useGetContest, useRouterContestSlug } from "./contest";

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
