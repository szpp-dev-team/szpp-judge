import { Icon, IconProps } from "@chakra-ui/react";

export type RoundedCrossIconProps = IconProps;

export const RoundedCrossIcon = (props: RoundedCrossIconProps) => (
  <Icon viewBox="0 0 16 16" fill="none" color="#CD0000" {...props}>
    <path
      d="M3.5 3.49998L12.5 12.5M12.5 3.49998L3.5 12.5"
      stroke="currentColor"
      strokeWidth="1.8"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </Icon>
);
