import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { Login } from "./Login";

export const LoginPage = () => {
  return (
    <WithHeaderFooter>
      <Head>
        <title>{pageTitle("ログイン")}</title>
      </Head>
      <Login />
    </WithHeaderFooter>
  );
};
