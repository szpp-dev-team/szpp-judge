import { WithHeaderFooter } from "@/src/components/layout/WithHeaderFooter";
import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { ContestList } from "./ContestList";

export const ContestListPage = () => {
  return (
    <WithHeaderFooter>
      <Head>
        <title>{pageTitle("コンテスト")}</title>
      </Head>
      <ContestList />
    </WithHeaderFooter>
  );
};
