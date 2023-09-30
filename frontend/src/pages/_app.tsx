import customizedTheme from "@/src/config/theme";
import { ChakraProvider } from "@chakra-ui/react";
import { TransportProvider } from "@connectrpc/connect-query";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import type { AppProps } from "next/app";
import { backendGrpcTransport } from "../config/grpc";

import("../mocks").catch((reason) => console.error(reason));

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
