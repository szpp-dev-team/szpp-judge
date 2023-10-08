import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { useRouter } from "next/router";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestSubmissionList } from "./ContestSubmissionList";

export const ContestSubmissionListPage = () => {
  const { query } = useRouter();
  const title = useContestPageTitle("me" in query ? "自分の提出" : "すべての提出");
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestSubmissionList />
    </ContestLayout>
  );
};
