import { ChevronDownIcon } from "@chakra-ui/icons";
import { Box, Divider, Flex, Icon, Link, Text, useOutsideClick } from "@chakra-ui/react";
import type { BoxProps, LinkProps } from "@chakra-ui/react";
import NextLink from "next/link";
import { useEffect, useRef, useState } from "react";
import type { IconType } from "react-icons";
import { IoCodeSlash, IoExit, IoPerson, IoSettings, IoTrophy } from "react-icons/io5";

type Color = BoxProps["bg"];

export type GlobalHeaderProps = {
  contestSlug?: string;
  contestTitle?: string;
  username?: string;
  behindFillColor?: Color;
};

export const GLOBAL_HEADER_H = "52px";

const BG_COLOR: Color = "teal.500";
const FG_COLOR: Color = "white";

const ITEM_COMMON_PROPS = {
  display: "flex",
  alignItems: "center",
  px: "0.75rem",
  h: "100%",
  fontSize: "16px",
  fontWeight: "semibold",
} as const satisfies BoxProps;

export const GlobalHeader = ({
  contestSlug,
  contestTitle,
  username,
  behindFillColor,
}: GlobalHeaderProps) => {
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
        bg={BG_COLOR}
        color={FG_COLOR}
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
          {username ? <UserMenu username={username} /> : (
            <>
              <NavItemLink href="/register" px="1.25rem">登録</NavItemLink>
              <NavItemLink href="/login" px="1.25rem">ログイン</NavItemLink>
            </>
          )}
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
        {...ITEM_COMMON_PROPS}
        _hover={{ background: "teal.600" }}
        {...props}
      >
        {children}
      </Link>
    </Box>
  );
};

const UserMenu = ({ username }: {
  username: string;
} & BoxProps) => {
  const [menuVisible, setMenuVisible] = useState(false);

  // transition のための maxH に指定する値。コンポーネントマウント時に menuListRef から値を取得する
  const [menuListHeight, setMenuListHeight] = useState(0);
  const menuListRef = useRef<HTMLUListElement>(null);
  useEffect(() => {
    setMenuListHeight(menuListRef.current?.getBoundingClientRect().height ?? 0);
  }, []);

  const menuWrapperRef = useRef<HTMLLIElement>(null);
  useOutsideClick({
    ref: menuWrapperRef,
    handler: () => {
      setMenuVisible(false);
    },
  });

  return (
    <li ref={menuWrapperRef} style={{ position: "relative" }}>
      <Flex
        position="relative"
        as="button"
        {...ITEM_COMMON_PROPS}
        zIndex={50}
        onClick={() => setMenuVisible(!menuVisible)}
      >
        <Text as="span">{username}</Text>
        <ChevronDownIcon ml={1} fontSize="18px" transform={`rotate(${menuVisible ? "180deg" : "0"})`} />
      </Flex>
      <Box
        position="absolute"
        zIndex={49}
        top={GLOBAL_HEADER_H}
        right={0}
        overflowY="hidden"
        transitionDuration="250ms"
        transitionProperty="max-height"
        maxH={menuVisible ? `${menuListHeight}px` : "0"}
      >
        <Box
          as="ul"
          ref={menuListRef}
          w="max-content"
          bg="gray.100"
          color="teal.900"
          borderLeft="1px"
          borderBottom="1px"
          borderColor="gray.400"
        >
          <MenuItem href={`/users/${username}`} icon={IoPerson}>マイプロフィール</MenuItem>
          <MenuItem href={`/users/${username}/contest-history`} icon={IoTrophy}>コンテスト参加履歴</MenuItem>
          <MenuItem href={`/users/${username}/submissions`} icon={IoCodeSlash}>提出一覧</MenuItem>
          <Divider />
          <MenuItem href={`/settings`} icon={IoSettings}>設定</MenuItem>
          <Divider />
          <MenuItem as="button" icon={IoExit}>ログアウト</MenuItem>
        </Box>
      </Box>
    </li>
  );
};

const MenuItem = ({ icon, children, ...props }: {
  icon: IconType;
} & LinkProps) => {
  return (
    <Box as="li" display="block">
      <Link
        as={NextLink}
        display="flex"
        alignItems="center"
        gap={2}
        py={2}
        px={4}
        fontSize="14px"
        w="100%"
        textAlign="left"
        _hover={{ background: "gray.200" }}
        {...props}
      >
        <Icon as={icon} />
        {children}
      </Link>
    </Box>
  );
};
