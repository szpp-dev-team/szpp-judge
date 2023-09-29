import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestLayout } from "../../layout/ContestLayout";
import { TaskCollection } from "./TaskCollection";

export const TaskCollectionPage = () => {
  const title = useContestPageTitle("問題");
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <TaskCollection />
    </ContestLayout>
  );
};
