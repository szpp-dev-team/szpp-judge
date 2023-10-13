import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { useRouter } from "next/router";
import { ContestContentGuard } from "../../guard/ContestContentGuard";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestSubmissionList } from "./ContestSubmissionList";

export const ContestSubmissionListPage = () => {
  const { query } = useRouter();
  const mode = "me" in query ? "me" : "all";
  const contentName = mode === "me" ? "自分の提出" : "すべての提出";
  const title = useContestPageTitle(contentName);
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestContentGuard contentName={contentName}>
        <ContestSubmissionList mode={mode} heading={contentName} />
      </ContestContentGuard>
    </ContestLayout>
  );
};
