import { useAuthUserValue } from "@/src/globalStates/credential";
import { Box, type BoxProps, Flex } from "@chakra-ui/react";
import { type ReactNode } from "react";
import { GlobalFooter, type GlobalFooterProps } from "../ui/GlobalFooter";
import { GLOBAL_HEADER_H, GlobalHeader, type GlobalHeaderProps } from "../ui/GlobalHeader";

export type WithHeaderFooterProps = {
  children: ReactNode;
  headerProps?: GlobalHeaderProps;
  footerProps?: GlobalFooterProps;
  leftChildren?: ReactNode;
  rightChildren?: ReactNode;
};

export const WithHeaderFooter = ({
  children,
  headerProps,
  footerProps,
  leftChildren,
  rightChildren,
}: WithHeaderFooterProps) => {
  const bg: BoxProps["bg"] = "gray.200";
  const user = useAuthUserValue();
  return (
    <>
      <GlobalHeader username={user?.username} {...headerProps} behindFillColor={bg} />
      <Flex minW="100%" bg={bg}>
        {leftChildren}
        {
          // 子要素で height=100% がうまく機能するためには、親要素で height が auto になってはならない
          // そこで、適当に height=1px を設定することで非 auto にするというトリックを施している
        }
        <Flex flexDirection="column" minH="100vh" h="1px" pt={GLOBAL_HEADER_H} flexGrow={1}>
          <Box as="main" bg={bg} flexGrow={1}>{children}</Box>
          <GlobalFooter {...footerProps} />
        </Flex>
        {rightChildren}
      </Flex>
    </>
  );
};
