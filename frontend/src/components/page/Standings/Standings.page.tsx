import { ContestLayout } from "@/src/components/layout/ContestLayout";
import { useContestPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { Standings } from "./Standings";

export const StandingsPage = () => {
  const title = useContestPageTitle("順位表");
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <Standings />
    </ContestLayout>
  );
};
