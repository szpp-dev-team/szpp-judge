import { backendTransportWithOnlyBaseUrl } from "@/src/config/connectRpc";
import { HealthcheckService } from "@/src/gen/proto/backend/v1/healthcheck_service-HealthcheckService_connectquery";
import { createPromiseClient } from "@connectrpc/connect";
import { Button, Container } from "@chakra-ui/react";
import type { NextPage } from "next";
import { useState } from "react";

const Page: NextPage = () => {
  const cli = createPromiseClient(
    HealthcheckService,
    backendTransportWithOnlyBaseUrl,
  );

  const [waitingPing, setWaitingPing] = useState(false);

  const handlePingButtonClick = async () => {
    console.log("requesting ping");
    setWaitingPing(true);
    try {
      const resp = await cli.ping({ name: "Bob" });
      console.log(resp);
    } finally {
      setWaitingPing(false);
    }
  };

  return (
    <Container>
      <header>Healthcheck Page</header>
      <Button
        colorScheme="teal"
        onClick={handlePingButtonClick}
        isLoading={waitingPing}
      >
        Ping
      </Button>
    </Container>
  );
};

export default Page;
