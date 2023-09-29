import { usePageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { Top } from "./Top";

export const TopPage = () => {
  const title = usePageTitle("トップ");
  return (
    <WithHeaderFooter>
      <Head>
        <title>{title}</title>
      </Head>
      <Top />
    </WithHeaderFooter>
  );
};
