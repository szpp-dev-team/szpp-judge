import { useContestTaskDetailPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestTaskDetail } from "./ContestTaskDetail";

export const ContestTaskDetailPage = () => {
  const title = useContestTaskDetailPageTitle();
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestTaskDetail />
    </ContestLayout>
  );
};
