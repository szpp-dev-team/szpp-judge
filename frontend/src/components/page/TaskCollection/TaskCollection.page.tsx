import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestContentGuard } from "../../guard/ContestContentGuard";
import { ContestLayout } from "../../layout/ContestLayout";
import { TaskCollection } from "./TaskCollection";

export const TaskCollectionPage = () => {
  const title = useContestPageTitle("問題");
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestContentGuard contentName="問題一覧">
        <TaskCollection />
      </ContestContentGuard>
    </ContestLayout>
  );
};
