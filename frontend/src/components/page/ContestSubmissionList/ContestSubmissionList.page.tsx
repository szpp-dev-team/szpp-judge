import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { useRouter } from "next/router";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestSubmissionList } from "./ContestSubmissionList";

export const ContestSubmissionListPage = () => {
  const { query } = useRouter();
  const mode = "me" in query ? "me" : "all";
  const heading = mode === "me" ? "自分の提出" : "すべての提出";
  const title = useContestPageTitle(heading);
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestSubmissionList mode={mode} heading={heading} />
    </ContestLayout>
  );
};
