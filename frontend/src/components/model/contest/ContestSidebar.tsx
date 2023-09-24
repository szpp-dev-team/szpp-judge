import type { ScoreStatus } from "@/src/model/task";
import { calcNthTaskSeq } from "@/src/usecases/contest";
import { Duration, fmtDatetime } from "@/src/util/time";
import { Box, FormControl, FormLabel, Icon, Link, Switch, Text } from "@chakra-ui/react";
import type { BoxProps, LinkProps } from "@chakra-ui/react";
import NextLink from "next/link";
import React, { type ReactNode, useMemo, useState } from "react";
import type { IconType } from "react-icons";
import { IoBarChart, IoChatboxEllipses, IoEarthSharp, IoHome, IoList, IoPerson, IoSchool } from "react-icons/io5";
import { SIDEBAR_TOGGLE_KNOB_H, SIDEBAR_TOGGLE_KNOB_TOP, SidebarToggleKnob } from "../../ui/SidebarToggleKnob";
import { ScoreStatusIcon } from "../task/ScoreStatusIcon";

export type ContestSidebarProps = {
  /** 画面上部から下へずらす量。GlobalHeaderの高さを期待する。"0px"など単位も必要。 */
  top?: string | "0px";
  /** コンテスト開始日時 */
  startAt?: Date;
  /** コンテスト終了日時 */
  endAt?: Date;
  /** 現在時刻 */
  now: Date;
  /** コンテストの slug */
  slug: string;
  /** コンテストの問題と得点状況 */
  tasks: Array<{
    id: number;
    title: string;
    scoreStatus?: ScoreStatus;
  }>;
} & Omit<BoxProps, "top" | "children">;

type Color = BoxProps["color"];
const FILL_COLOR: Color = "gray.100";
const FONT_COLOR: Color = "teal.900";
const DIVIDER_COLOR: Color = "gray.300";
const BORDER_COLOR: Color = "gray.400";

const SIDEBAR_MAIN_PANE_WIDTH = "14rem";

export const ContestSidebar = ({ top = "0px", ...props }: ContestSidebarProps) => {
  const [fixedShow, setFixedShow] = useState(true);
  const [temporaryShow, setTemporaryShow] = useState(false);

  return (
    <Box
      as="aside"
      minW={fixedShow ? SIDEBAR_MAIN_PANE_WIDTH : 0}
    >
      <Box
        as="nav"
        position="fixed"
        zIndex={45}
        top={top}
        left={0}
        h={`calc(100vh - ${top})`}
        maxH={`calc(100vh - ${top})`}
      >
        <SidebarMainPane
          zIndex={47}
          {...props}
          temporaryShow={temporaryShow}
          fixedShow={fixedShow}
          onToggleKnobClick={() => {
            if (!fixedShow && temporaryShow) {
              setTemporaryShow(false);
            } else {
              setFixedShow(!fixedShow);
            }
          }}
          onFixSwitchChange={() => {
            setTemporaryShow(true);
            setFixedShow(!fixedShow);
          }}
          onMouseLeave={() => setTemporaryShow(false)}
        />
        <SidebarHoverShowArea
          zIndex={46}
          onMouseEnter={() => setTemporaryShow(true)}
        />
      </Box>
    </Box>
  );
};

const SidebarHoverShowArea = (props: BoxProps) => {
  const w = "60px";
  return (
    <Box
      position="absolute"
      left={0}
      top={`calc(${SIDEBAR_TOGGLE_KNOB_TOP} + ${SIDEBAR_TOGGLE_KNOB_H})`}
      h={`calc(100% - ${SIDEBAR_TOGGLE_KNOB_TOP} - ${SIDEBAR_TOGGLE_KNOB_H})`}
      w={w}
      {...props}
    >
    </Box>
  );
};

type SidebarMainPaneProps = {
  temporaryShow: boolean;
  fixedShow: boolean;
  onToggleKnobClick?: React.MouseEventHandler<HTMLButtonElement>;
  onFixSwitchChange?: React.ChangeEventHandler<HTMLInputElement>;
} & Omit<ContestSidebarProps, "top" | "hidden">;

const SidebarMainPane = ({
  temporaryShow,
  fixedShow,
  onToggleKnobClick,
  onFixSwitchChange,
  startAt,
  endAt,
  now,
  slug,
  tasks,
  ...props
}: SidebarMainPaneProps) => {
  const contestRootPath = `/contests/${slug}`;
  const contestStarted = startAt ? now >= startAt : false;
  const contestFinished = endAt ? now >= endAt : false;

  return (
    <Box
      position="relative"
      h="100%"
      w={SIDEBAR_MAIN_PANE_WIDTH}
      minW={SIDEBAR_MAIN_PANE_WIDTH}
      py={2}
      transform={`translateX(${fixedShow || temporaryShow ? 0 : `-${SIDEBAR_MAIN_PANE_WIDTH}`})`}
      transition="transform"
      transitionDuration="200ms"
      display="flex"
      flexDirection="column"
      bg={FILL_COLOR}
      color={FONT_COLOR}
      borderRight="1px"
      borderColor={BORDER_COLOR}
      shadow="lg"
      fontSize="sm"
      {...props}
    >
      <FormControl display="flex" justifyContent="center" alignItems="center" my={1}>
        <FormLabel fontSize="xs" mb={0} cursor="pointer">サイドバーを常に表示する</FormLabel>
        <Switch isChecked={fixedShow} size="sm" onChange={onFixSwitchChange} />
      </FormControl>
      <SidebarToggleKnob
        bg={FILL_COLOR}
        color="teal.700"
        borderColor={BORDER_COLOR}
        iconDirection={fixedShow || temporaryShow ? "hide" : "show"}
        onClick={onToggleKnobClick}
      />
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
              borderColor={DIVIDER_COLOR}
              py={2}
              my={1}
            >
              <Box as="ul" listStyleType="none" overflowY="auto" overscrollBehavior="contain">
                {tasks.map((t, i) => (
                  <SidebarLinkItem
                    leftElem={
                      <>
                        <ScoreStatusIcon status={t.scoreStatus} mr={1} />
                        <TaskSeqBadge seq={calcNthTaskSeq(i, tasks.length)} />
                      </>
                    }
                    key={t.id}
                    text={t.title}
                    href={`${contestRootPath}/tasks/${i + 1}`}
                    fontSize="xs"
                    pl={2}
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
      <Box
        as="ul"
        listStyleType="none"
        mt="auto"
        mx={4}
        py={2}
        textAlign="center"
        borderTop="1px"
        borderColor={DIVIDER_COLOR}
      >
        {(startAt == null || endAt == null) ? <></> : (
          <>
            <SidebarRemainingTime startAt={startAt} endAt={endAt} now={now} />
            <SidebarDatetime label="開始" datetime={startAt} />
            <SidebarDatetime label="終了" datetime={endAt} />
          </>
        )}
      </Box>
    </Box>
  );
};

const TaskSeqBadge = ({ seq }: { seq: string }) => {
  return (
    <Text
      as="span"
      display="block"
      minW="1.75em"
      textAlign="center"
      mx={1}
      px={1}
      fontWeight="semibold"
      bg="blackAlpha.100"
      borderRadius={4}
    >
      {seq}
    </Text>
  );
};

const SidebarLinkItem = ({ icon, leftElem, text, href, ...props }: {
  icon?: IconType;
  leftElem?: ReactNode;
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
        {leftElem}
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
      <Text as="span" fontSize="2xs">{label}</Text>
      <Text as="time" dateTime={remainTime} display="block" fontSize="md" fontWeight="medium">{remainTime}</Text>
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
    <Box as="li" fontSize="xs">
      <Text as="span" mr={2}>{label}</Text>
      <Text as="time" dateTime={machineFmt}>{humanFmt}</Text>
    </Box>
  );
};
