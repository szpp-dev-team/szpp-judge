import { Box, Flex, Link } from "@chakra-ui/react";
import type { BoxProps, LinkProps } from "@chakra-ui/react";
import NextLink from "next/link";

export type GlobalHeaderProps = {
  contestSlug?: string;
  contestTitle?: string;
  behindFillColor?: BoxProps["bg"];
};

export const GLOBAL_HEADER_H = "52px";

export const GlobalHeader = ({ contestSlug, contestTitle, behindFillColor }: GlobalHeaderProps) => {
  return (
    <Box as="header" w="100%" bg={behindFillColor}>
      <Flex
        as="nav"
        justifyContent="space-between"
        position="fixed"
        top="0"
        left="0"
        zIndex={50}
        h={GLOBAL_HEADER_H}
        w="100%"
        bg="teal.500"
        color="white"
      >
        <Flex as="ul" listStyleType="none">
          <NavItemLink href="/" px="1.25rem">SZPP Judge</NavItemLink>
          {contestSlug ? <NavItemLink href={`/contests/${contestSlug}`}>{contestTitle}</NavItemLink> : (
            <>
              <NavItemLink href="/contests">コンテスト</NavItemLink>
              <NavItemLink href="/tasks">問題</NavItemLink>
            </>
          )}
        </Flex>
        <Flex as="ul" listStyleType="none">
          <NavItemLink href="/register" px="1.25rem">登録</NavItemLink>
          <NavItemLink href="/login" px="1.25rem">ログイン</NavItemLink>
        </Flex>
      </Flex>
    </Box>
  );
};

const NavItemLink = ({ children, ...props }: LinkProps) => {
  return (
    <Box as="li" display="block">
      <Link
        as={NextLink}
        display="flex"
        alignItems="center"
        px="0.75rem"
        h="100%"
        fontSize="16px"
        fontWeight="semibold"
        _hover={{ background: "teal.600" }}
        {...props}
      >
        {children}
      </Link>
    </Box>
  );
};
