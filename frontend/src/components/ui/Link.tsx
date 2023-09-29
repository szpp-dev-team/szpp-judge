import { Link as ChakraLink } from "@chakra-ui/react";
import type { LinkProps as ChakraLinkProps } from "@chakra-ui/react";
import NextLink from "next/link";

export type LinkProps = ChakraLinkProps;

export const Link = ({ children, ...props }: LinkProps) => {
  return (
    <ChakraLink as={NextLink} color="blue.600" {...props}>
      {children}
    </ChakraLink>
  );
};
