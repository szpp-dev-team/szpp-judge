import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { ContestLayout } from "../../layout/ContestLayout";
import { TaskCollection } from "./TaskCollection";

export const TaskCollectionPage = () => {
  return (
    <ContestLayout>
      <Head>
        <title>{pageTitle("問題")}</title>
      </Head>
      <TaskCollection />
    </ContestLayout>
  );
};
