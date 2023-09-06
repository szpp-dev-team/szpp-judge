import { Box, Flex } from "@chakra-ui/react";
import { type FC, ReactNode } from "react";
import { GlobalFooter } from "../ui/GlobalFooter";
import { GlobalHeader } from "../ui/GlobalHeader";

export type WithHeaderFooterProps = {
  children: ReactNode;
};

export const WithHeaderFooter: FC<WithHeaderFooterProps> = ({ children }) => {
  // 子要素で height=100% がうまく機能するためには、親要素で height が auto になってはならない
  // そこで、適当に height=1px を設定することで非 auto にするというトリックを施している
  return (
    <Flex flexDirection="column" minH="100vh" h="1px" bg="gray.200">
      <GlobalHeader />
      <Box as="main" height="100%">{children}</Box>
      <GlobalFooter mt="auto" />
    </Flex>
  );
};
