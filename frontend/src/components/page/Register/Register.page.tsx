import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { Register } from "./Register";

export const RegisterPage = () => {
  return (
    <WithHeaderFooter>
      <Head>
        <title>{pageTitle("ユーザ登録")}</title>
      </Head>
      <Register />
    </WithHeaderFooter>
  );
};
