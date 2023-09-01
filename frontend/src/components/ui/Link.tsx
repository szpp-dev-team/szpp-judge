import { Link as ChakraLink } from "@chakra-ui/react";
import type { LinkProps as ChakraLinkProps } from "@chakra-ui/react";
import NextLink from "next/link";
import type { FC } from "react";

export type LinkProps = ChakraLinkProps;

export const Link: FC<ChakraLinkProps> = ({ children, ...props }) => {
  return (
    <ChakraLink as={NextLink} color="blue.500" {...props}>
      {children}
    </ChakraLink>
  );
};
