import { Difficulty } from "@/src/model/task";
import { captalize } from "@/src/util/string";
import { Box, type BoxProps } from "@chakra-ui/react";
import type { FC } from "react";

export type DifficultyBadgeProps = {
  dif: Difficulty;
} & Omit<BoxProps, "children">;

const Base = ({ d, ...props }: {
  d: Difficulty;
} & Omit<BoxProps, "children">) => {
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
      {...props}
    >
      {captalize(d)}
    </Box>
  );
};

export const DifficultyBadge: FC<DifficultyBadgeProps> = ({ dif, ...props }) => {
  switch (dif) {
    case "beginner":
      return <Base d={dif} bg="gray.200" borderColor="gray.400" color="gray.700" {...props} />;
    case "easy":
      return <Base d={dif} bg="green.100" borderColor="green.500" color="green.800" {...props} />;
    case "medium":
      return <Base d={dif} bg="cyan.100" borderColor="cyan.500" color="cyan.900" {...props} />;
    case "hard":
      return <Base d={dif} bg="orange.100" borderColor="orange.500" color="orange.700" {...props} />;
    case "impossible":
      return <Base d={dif} bg="black" borderColor="#E56E6E" color="#FF7C7C" fontWeight="bold" {...props} />;
    default: {
      const exhaustiveCheck: never = dif;
      throw new Error(`Invlid difficulty: ${exhaustiveCheck}`);
    }
  }
};
