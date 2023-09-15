import { ChevronLeftIcon } from "@chakra-ui/icons";
import { Box, HTMLChakraProps, IconProps } from "@chakra-ui/react";

export type SidebarToggleKnobProps = {
  iconDirection?: "hide" | "show";
  iconProps?: IconProps;
} & Omit<HTMLChakraProps<"button">, "children" | "top" | "bottom" | "left" | "right" | "height" | "width" | "h" | "w">;

export const SIDEBAR_TOGGLE_KNOB_TOP = "40px";
export const SIDEBAR_TOGGLE_KNOB_H = "48px";
export const SIDEBAR_TOGGLE_KNOB_W = "28px";

export const SidebarToggleKnob = ({ iconDirection = "hide", iconProps, ...props }: SidebarToggleKnobProps) => {
  return (
    <Box
      as="button"
      position="absolute"
      top={SIDEBAR_TOGGLE_KNOB_TOP}
      right={`-${SIDEBAR_TOGGLE_KNOB_W}`}
      h={SIDEBAR_TOGGLE_KNOB_H}
      w={SIDEBAR_TOGGLE_KNOB_W}
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
