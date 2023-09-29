import { WithHeaderFooter } from "@/src/components/layout/WithHeaderFooter";
import { usePageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestList } from "./ContestList";

export const ContestListPage = () => {
  const title = usePageTitle("コンテスト一覧");
  return (
    <WithHeaderFooter>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestList />
    </WithHeaderFooter>
  );
};
