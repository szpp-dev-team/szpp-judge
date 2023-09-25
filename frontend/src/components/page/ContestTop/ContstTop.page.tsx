import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestLayout } from "../../layout/ContestLayout";
import { ContestTop } from "./ContestTop";

export const ContestTopPage = () => {
  const title = useContestPageTitle("トップ");
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestTop />
    </ContestLayout>
  );
};
