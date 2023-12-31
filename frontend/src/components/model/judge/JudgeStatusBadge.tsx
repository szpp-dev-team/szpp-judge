import { JudgeStatus, type JudgeTestcaseProgress } from "@/src/model/judge";
import { Box, type BoxProps } from "@chakra-ui/react";
import type { FC } from "react";

export type JudgeStatusBadgeProps = {
  /**
   * ジャッジの状態
   */
  status: JudgeStatus;
  /**
   * true なら英語の省略表記 (AC, WA, etc.) で表示する
   */
  abbr?: boolean;
  /**
   * ジャッジの進捗を、完了したテストケースの個数とテストケースの総数を基準に表示する
   */
  progress?: JudgeTestcaseProgress;
} & Omit<BoxProps, "children">;

export const JudgeStatusBadge: FC<JudgeStatusBadgeProps> = ({ status, abbr, progress, ...props }) => {
  const bg = ((): BoxProps["bg"] => {
    switch (status) {
      case "WJ":
        return "gray.500";
      case "AC":
        return "#59C163";
      case "IE":
        return "purple.500";
      default:
        return "#E4A03E";
    }
  })();
  const text = ((): string => {
    if (progress) {
      const { done, total } = progress;
      return `${done}/${total} ...`;
    }
    return (abbr ? JudgeStatus.toEnglishAbbr : JudgeStatus.toJapanese)(status);
  })();
  return (
    <Box
      display="inline-flex"
      justifyContent="center"
      alignItems="center"
      w={abbr ? "5.5rem" : "7rem"}
      fontSize={abbr || progress ? "sm" : "xs"}
      py="0.2rem"
      color="white"
      bg={bg}
      borderRadius="6px"
      fontWeight="semibold"
      {...props}
    >
      {text}
    </Box>
  );
};
