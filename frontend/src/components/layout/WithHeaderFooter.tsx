import { Flex } from "@chakra-ui/react";
import { FC, ReactNode } from "react";
import { GlobalFooter } from "../ui/GlobalFooter";
import { GlobalHeader } from "../ui/GlobalHeader";

export type WithHeaderFooterProps = {
  children: ReactNode;
};

export const WithHeaderFooter: FC<WithHeaderFooterProps> = ({ children }) => {
  return (
    <Flex flexDirection="column" minHeight="100vh" bg="gray.200">
      <GlobalHeader />
      <main>{children}</main>
      <GlobalFooter mt="auto" />
    </Flex>
  );
};
