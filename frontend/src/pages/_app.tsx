import customizedTheme from "@/src/config/theme";
import { defaultErrorMap } from "@/src/config/zod";
import { ChakraProvider } from "@chakra-ui/react";
import { TransportProvider } from "@connectrpc/connect-query";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import type { AppProps } from "next/app";
import { z } from "zod";
import { backendGrpcTransport } from "../config/grpc";

import("../mocks").catch((reason) => console.error(reason));

z.setErrorMap(defaultErrorMap);

const queryClient = new QueryClient();

const MyApp = ({ Component, pageProps }: AppProps) => {
  return (
    <TransportProvider transport={backendGrpcTransport}>
      <QueryClientProvider client={queryClient}>
        <ChakraProvider toastOptions={{ defaultOptions: { isClosable: true } }} theme={customizedTheme}>
          <Component {...pageProps} />
        </ChakraProvider>
      </QueryClientProvider>
    </TransportProvider>
  );
};

export default MyApp;
