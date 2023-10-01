import customizedTheme from "@/src/config/theme";
import { ChakraProvider } from "@chakra-ui/react";
import { TransportProvider } from "@connectrpc/connect-query";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import type { AppProps } from "next/app";
import { createBackendGrpcTransport } from "../config/grpc";
import { useAccessTokenClaimValue, useCredentialValue } from "../globalStates/credential";
import { useRefreshAccessTokenWithoutQueryClient } from "../usecases/auth";

if (["1", "true", "yes"].includes(process.env.MOCK_ENABLED || "")) {
  import("../mocks").catch((reason) => console.error(reason));
}

const queryClient = new QueryClient();

const MyApp = ({ Component, pageProps }: AppProps) => {
  const cred = useCredentialValue();
  const claim = useAccessTokenClaimValue();
  const { refreshAccessToken } = useRefreshAccessTokenWithoutQueryClient();
  const transport = createBackendGrpcTransport({
    cred,
    accessTokenExpireAt: claim?.exp,
    refreshAccessToken: async () => (await refreshAccessToken()).accessToken,
  });

  return (
    <TransportProvider transport={transport}>
      <QueryClientProvider client={queryClient}>
        <ChakraProvider toastOptions={{ defaultOptions: { isClosable: true } }} theme={customizedTheme}>
          <Component {...pageProps} />
        </ChakraProvider>
      </QueryClientProvider>
    </TransportProvider>
  );
};

export default MyApp;
