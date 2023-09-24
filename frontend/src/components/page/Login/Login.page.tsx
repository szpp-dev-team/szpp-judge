import { usePageTitle } from "@/src/usecases/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { Login } from "./Login";

export const LoginPage = () => {
  const title = usePageTitle("ログイン");
  return (
    <WithHeaderFooter>
      <Head>
        <title>{title}</title>
      </Head>
      <Login />
    </WithHeaderFooter>
  );
};
