import { Box, type BoxProps, Flex } from "@chakra-ui/react";
import { type ReactNode } from "react";
import { GlobalFooter, type GlobalFooterProps } from "../ui/GlobalFooter";
import { GlobalHeader, type GlobalHeaderProps } from "../ui/GlobalHeader";

export type WithHeaderFooterProps = {
  children: ReactNode;
  headerProps?: GlobalHeaderProps;
  bodyProps?: BoxProps;
  footerProps?: GlobalFooterProps;
};

export const WithHeaderFooter = ({
  children,
  headerProps,
  bodyProps,
  footerProps,
}: WithHeaderFooterProps) => {
  // 子要素で height=100% がうまく機能するためには、親要素で height が auto になってはならない
  // そこで、適当に height=1px を設定することで非 auto にするというトリックを施している
  return (
    <Flex flexDirection="column" minH="100vh" h="1px" w="fit-content" minW="100%">
      <GlobalHeader {...headerProps} />
      <Box as="main" flexGrow={1} bg="gray.200" {...bodyProps}>{children}</Box>
      <GlobalFooter {...footerProps} />
    </Flex>
  );
};
