import { useSubmissionDetailPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestLayout } from "../../layout/ContestLayout";
import { SubmissionDetail } from "./SubmissionDetail";

export const SubmissionDetailPage = () => {
  const title = useSubmissionDetailPageTitle();
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <SubmissionDetail />
    </ContestLayout>
  );
};
