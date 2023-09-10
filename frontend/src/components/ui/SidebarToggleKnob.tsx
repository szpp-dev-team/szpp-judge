import { ChevronLeftIcon } from "@chakra-ui/icons";
import { Box, HTMLChakraProps, IconProps } from "@chakra-ui/react";

export type SidebarToggleKnobProps = {
  iconDirection?: "hide" | "show";
  iconProps?: IconProps;
} & Omit<HTMLChakraProps<"button">, "children">;

export const SidebarToggleKnob = ({ iconDirection = "hide", iconProps, ...props }: SidebarToggleKnobProps) => {
  const w = "28px";
  return (
    <Box
      as="button"
      position="absolute"
      top="40px"
      right={`-${w}`}
      h="48px"
      w={w}
      bg="gray.100"
      color="teal.700"
      border="1px"
      borderColor="gray.400"
      borderTopRightRadius="6px"
      borderBottomRightRadius="6px"
      {...props}
    >
      <ChevronLeftIcon
        fontSize="24px"
        transform="auto-gpu"
        rotate={iconDirection === "show" ? "180deg" : 0}
        {...iconProps}
      />
    </Box>
  );
};
