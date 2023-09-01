import { HealthcheckService } from "@/src/gen/proto/backend/v1/services_connect";
import { backendGrpcTransport } from "@/src/config/grpc";
import { createPromiseClient } from "@bufbuild/connect";
import { Button, Container } from "@chakra-ui/react";
import type { NextPage } from "next";
import { useState } from "react";

const Page: NextPage = () => {
  const cli = createPromiseClient(
    HealthcheckService,
    backendGrpcTransport
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
