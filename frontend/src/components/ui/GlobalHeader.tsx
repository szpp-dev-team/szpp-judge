import { Box, Flex, Link } from "@chakra-ui/react";
import type { LinkProps } from "@chakra-ui/react";
import NextLink from "next/link";
import type { FC } from "react";

export type GlobalHeaderProps = {
  contestID?: number;
  contestTitle?: string;
};

const NavItemLink: FC<LinkProps> = ({ children, ...props }) => {
  return (
    <Link
      as={NextLink}
      display="block"
      py="1.25rem"
      px="0.75rem"
      fontSize="md"
      fontWeight="semibold"
      _hover={{ background: "teal.600" }}
      {...props}
    >
      {children}
    </Link>
  );
};

export const GlobalHeader: FC<GlobalHeaderProps> = ({ contestID: contestId, contestTitle }) => {
  return (
    <Box as="header" position="sticky" width="100%" top="0" left="0" zIndex={50}>
      <Flex as="nav" justifyContent="space-between" bg="teal.500" color="white">
        <Flex>
          <NavItemLink href="/" px="1.25rem">SZPP Judge</NavItemLink>
          {contestId ? <NavItemLink href={`/contests/${contestId}`}>{contestTitle}</NavItemLink> : (
            <>
              <NavItemLink href="/contests">コンテスト</NavItemLink>
              <NavItemLink href="/tasks">問題</NavItemLink>
            </>
          )}
        </Flex>
        <Flex>
          <NavItemLink href="/register" px="1.25rem">登録</NavItemLink>
          <NavItemLink href="/login" px="1.25rem">ログイン</NavItemLink>
        </Flex>
      </Flex>
    </Box>
  );
};
