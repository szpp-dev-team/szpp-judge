import type { Difficulty } from "@/src/model/task";
import { capitalize } from "@/src/util/string";
import { Box, type BoxProps } from "@chakra-ui/react";
import type { FC } from "react";

export type DifficultyBadgeProps = {
  dif: Difficulty;
} & Omit<BoxProps, "children">;

export const DifficultyBadge: FC<DifficultyBadgeProps> = ({ dif, ...props }) => {
  const p = ((): BoxProps => {
    switch (dif) {
      case "beginner":
        return { bg: "gray.200", borderColor: "gray.400", color: "gray.700" };
      case "easy":
        return { bg: "green.100", borderColor: "green.500", color: "green.800" };
      case "medium":
        return { bg: "cyan.100", borderColor: "cyan.500", color: "cyan.900" };
      case "hard":
        return { bg: "orange.100", borderColor: "orange.500", color: "orange.700" };
      case "impossible":
        return { bg: "black", borderColor: "#E56E6E", color: "#FF7C7C", fontWeight: "bold" };
      default: {
        const exhaustiveCheck: never = dif;
        throw new Error(`Invlid difficulty: ${exhaustiveCheck}`);
      }
    }
  })();

  return (
    <Box
      borderRadius="999px"
      display="inline-flex"
      justifyContent="center"
      alignItems="center"
      w="5.5rem"
      py="0.15rem"
      px="1rem"
      border="1px"
      borderStyle="solid"
      fontSize="xs"
      {...p}
      {...props}
    >
      {capitalize(dif)}
    </Box>
  );
};
