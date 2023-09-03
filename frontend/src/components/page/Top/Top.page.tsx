import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { Top } from "./Top";

export const TopPage = () => {
  return (
    <WithHeaderFooter>
      <Head>
        <title>{pageTitle("トップ")}</title>
      </Head>
      <Top />
    </WithHeaderFooter>
  );
};
