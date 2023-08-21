import type { NextPage } from "next";
import { useState } from "react";
import { Button, Container } from "@chakra-ui/react";
import { createGrpcWebTransport } from "@bufbuild/connect-web";
import { createPromiseClient } from "@bufbuild/connect";
import { HealthcheckService } from "@/src/gen/proto/backend/v1/services_connect";

const Page: NextPage = () => {
  const cli = createPromiseClient(
    HealthcheckService,
    createGrpcWebTransport({ baseUrl: "http://localhost:8080" })
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