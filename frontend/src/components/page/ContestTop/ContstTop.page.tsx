import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestTop } from "./ContestTop";

export const ContestTopPage = () => {
  return (
    <ContestLayout>
      <Head>
        <title>{pageTitle("コンテストトップ")}</title>
      </Head>
      <ContestTop />
    </ContestLayout>
  );
};
