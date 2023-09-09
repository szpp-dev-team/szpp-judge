import { Duration, fmtDatetime } from "@/src/util/time";
import { Box, BoxProps, Icon, Link, LinkProps, Text } from "@chakra-ui/react";
import NextLink from "next/link";
import { useMemo } from "react";
import type { IconType } from "react-icons";
import { IoBarChart, IoChatboxEllipses, IoEarthSharp, IoHome, IoList, IoPerson, IoSchool } from "react-icons/io5";

export type ContestSidebarProps = {
  /** 画面上部から下へずらす量。GlobalHeaderの高さを期待する。"0px"など単位も必要。 */
  top?: string | "0px";
  /** コンテスト開始日時 */
  startAt: Date;
  /** コンテスト終了日時 */
  endAt: Date;
  /** 現在時刻 */
  now: Date;
  /** コンテストの slug */
  slug: string;
  /** コンテストの問題と提出結果 */
  tasks: ReadonlyArray<{
    id: number;
    title: string;
    status?: "accepted" | "partialAccepted" | "rejected";
  }>;
} & Omit<BoxProps, "top" | "children">;

export const ContestSidebar = ({
  top = "0px",
  startAt,
  endAt,
  now,
  slug,
  tasks,
  ...props
}: ContestSidebarProps) => {
  const contestRootPath = `/contests/${slug}`;
  const contestStarted = now >= startAt;
  const contestFinished = now >= endAt;
  const width = "17rem";

  return (
    <Box
      as="aside"
      position="sticky"
      zIndex={45}
      top={top}
      left={0}
      h={`calc(100vh - ${top})`}
      maxH={`calc(100vh - ${top})`}
      w={width}
      minW={width}
      py={2}
      display="flex"
      flexDirection="column"
      justifyContent="space-between"
      bg="gray.100"
      color="teal.900"
      borderRight="1px"
      borderColor="gray.400"
      shadow="lg"
      fontSize="sm"
      {...props}
    >
      <Box as="ul" listStyleType="none" whiteSpace="nowrap" overflowY="auto" display="flex" flexDirection="column">
        <SidebarLinkItem text="コンテストトップ" icon={IoHome} href={`${contestRootPath}`} />
        {contestStarted && (
          <>
            <SidebarLinkItem text="質問" icon={IoChatboxEllipses} href={`${contestRootPath}/clarifications`} />
            <SidebarLinkItem text="問題" icon={IoList} href={`${contestRootPath}/tasks`} />
            <Box
              as="li"
              overflowY="auto"
              borderTop="1px"
              borderBottom="1px"
              borderColor="gray.300"
              py={2}
              my={1}
            >
              <Box as="ul" listStyleType="none" overflowY="auto" overscrollBehavior="contain">
                {tasks.map((t, i) => (
                  <SidebarLinkItem
                    key={t.id}
                    text={String.fromCharCode("A".charCodeAt(0) + i) + " - " + t.title}
                    href={`${contestRootPath}/tasks/${i + 1}`}
                    overflow="hidden"
                    textOverflow="ellipsis"
                  />
                ))}
              </Box>
            </Box>
            <SidebarLinkItem text="自分の提出結果" icon={IoPerson} href={`${contestRootPath}/submissions?me`} />
            {contestFinished && (
              <SidebarLinkItem text="すべての提出結果" icon={IoEarthSharp} href={`${contestRootPath}/submissions`} />
            )}
          </>
        )}
        <SidebarLinkItem text="順位表" icon={IoBarChart} href={`${contestRootPath}/standings`} />
        {contestFinished && <SidebarLinkItem text="解説" icon={IoSchool} href={`${contestRootPath}/editorial`} />}
      </Box>
      <Box as="ul" listStyleType="none" mx={4} py={2} textAlign="center" borderTop="1px" borderColor="gray.300">
        <SidebarRemainingTime startAt={startAt} endAt={endAt} now={now} />
        <SidebarDatetime label="開始" datetime={startAt} />
        <SidebarDatetime label="終了" datetime={endAt} />
      </Box>
    </Box>
  );
};

const SidebarLinkItem = ({ icon, text, href, ...props }: {
  icon?: IconType;
  href: string;
  text: string;
} & Omit<LinkProps, "children">) => {
  return (
    <Box as="li" display="block">
      <Link
        as={NextLink}
        px={4}
        py={2}
        href={href}
        display="flex"
        alignItems="center"
        _hover={{ bg: "gray.200" }}
        {...props}
      >
        {icon && <Icon as={icon} mr={2} />}
        <Text as="span" overflow="hidden" textOverflow="ellipsis">{text}</Text>
      </Link>
    </Box>
  );
};

const SidebarRemainingTime = ({ startAt, endAt, now }: {
  startAt: Date;
  endAt: Date;
  now: Date;
}): JSX.Element | undefined => {
  const render = (label: string, remainTime: string) => (
    <Box as="li">
      <Text as="span" fontSize="xs">{label}</Text>
      <Text as="time" dateTime={remainTime} display="block" fontSize="lg" fontWeight="medium">{remainTime}</Text>
    </Box>
  );
  if (now < startAt) {
    const untilStart = new Duration(now, startAt);
    return (untilStart.hours() < 24) ? render("コンテスト開始まで", untilStart.fmtHMS()) : undefined;
  }
  if (now < endAt) {
    const untilEnd = new Duration(now, endAt);
    return render("コンテスト終了まで", untilEnd.fmtHMS());
  }
  return <Text as="li" fontSize="xs" py={1}>このコンテストは終了しました</Text>;
};

const SidebarDatetime = ({ label, datetime }: {
  label: string;
  datetime: Date;
}) => {
  const machineFmt = useMemo(
    () => datetime.toISOString(),
    [datetime],
  );
  const humanFmt = useMemo(
    () => fmtDatetime(datetime),
    [datetime],
  );
  return (
    <Box as="li">
      <Text as="span" fontSize="xs" mr={2}>{label}</Text>
      <Text as="time" dateTime={machineFmt}>{humanFmt}</Text>
    </Box>
  );
};
