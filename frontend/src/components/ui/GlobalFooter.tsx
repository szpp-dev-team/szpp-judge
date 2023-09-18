import { Flex, Text } from "@chakra-ui/react";
import type { BoxProps } from "@chakra-ui/react";

export type GlobalFooterProps = Omit<BoxProps, "children">;

export const GlobalFooter = ({ ...props }: GlobalFooterProps) => {
  return (
    <Flex
      as="footer"
      justifyContent="center"
      alignItems="center"
      bg="gray.700"
      py="1.75rem"
      {...props}
    >
      <Text as="small" color="gray.300">
        © 静岡大学プログラミングサークル SZPP
      </Text>
    </Flex>
  );
};
