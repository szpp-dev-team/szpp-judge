import customizedTheme from "@/src/config/theme";
import { defaultErrorMap } from "@/src/config/zod";
import { ChakraProvider } from "@chakra-ui/react";
import type { AppProps } from "next/app";
import { z } from "zod";

import("../mocks").catch((reason) => console.error(reason));

z.setErrorMap(defaultErrorMap);

const MyApp = ({ Component, pageProps }: AppProps) => {
  return (
    <ChakraProvider toastOptions={{ defaultOptions: { isClosable: true } }} theme={customizedTheme}>
      <Component {...pageProps} />
    </ChakraProvider>
  );
};

export default MyApp;
