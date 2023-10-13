import { useContestTaskDetailPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestContentGuard } from "../../guard/ContestContentGuard";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestTaskDetail } from "./ContestTaskDetail";

export const ContestTaskDetailPage = () => {
  const title = useContestTaskDetailPageTitle();
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestContentGuard contentName="問題文">
        <ContestTaskDetail />
      </ContestContentGuard>
    </ContestLayout>
  );
};
