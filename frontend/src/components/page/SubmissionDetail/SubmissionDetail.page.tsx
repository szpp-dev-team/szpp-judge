import { useSubmissionDetailPageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { ContestContentGuard } from "../../guard/ContestContentGuard";
import { ContestLayout } from "../../layout/ContestLayout";
import { SubmissionDetail } from "./SubmissionDetail";

export const SubmissionDetailPage = () => {
  const title = useSubmissionDetailPageTitle();
  return (
    <ContestLayout>
      <Head>
        <title>{title}</title>
      </Head>
      <ContestContentGuard contentName="提出結果">
        <SubmissionDetail />
      </ContestContentGuard>
    </ContestLayout>
  );
};
