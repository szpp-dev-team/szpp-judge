import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { TaskCollection } from "./TaskCollection";

export const TaskCollectionPage = () => {
  return (
    <WithHeaderFooter>
      <Head>
        <title>{pageTitle("問題")}</title>
      </Head>
      <TaskCollection />
    </WithHeaderFooter>
  );
};
