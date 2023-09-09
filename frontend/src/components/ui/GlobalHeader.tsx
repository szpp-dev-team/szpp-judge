import { Box, Flex, Link } from "@chakra-ui/react";
import type { LinkProps } from "@chakra-ui/react";
import NextLink from "next/link";

export type GlobalHeaderProps = {
  contestSlug?: string;
  contestTitle?: string;
};

const NavItemLink = ({ children, ...props }: LinkProps) => {
  return (
    <Box as="li" display="block">
      <Link
        as={NextLink}
        display="block"
        py="1rem"
        px="0.75rem"
        fontSize="sm"
        fontWeight="semibold"
        _hover={{ background: "teal.600" }}
        {...props}
      >
        {children}
      </Link>
    </Box>
  );
};

export const GlobalHeader = ({ contestSlug, contestTitle }: GlobalHeaderProps) => {
  return (
    <Box as="header" position="sticky" width="100%" top="0" left="0" zIndex={50}>
      <Flex as="nav" justifyContent="space-between" bg="teal.500" color="white">
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
