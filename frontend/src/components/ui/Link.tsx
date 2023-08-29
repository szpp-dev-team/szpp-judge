import { Link as ChakraLink, LinkProps as ChakraLinkProps } from "@chakra-ui/react";
import NextLink from "next/link";
import { FC } from "react";

export type LinkProps = ChakraLinkProps;

export const Link: FC<ChakraLinkProps> = ({ children, ...props }) => {
  return (
    <ChakraLink as={NextLink} color="blue.500" {...props}>
      {children}
    </ChakraLink>
  );
};
