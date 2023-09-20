import { Icon, IconProps } from "@chakra-ui/react";

export type RoundedCheckIconProps = IconProps;

export const RoundedCheckIcon = (props: RoundedCheckIconProps) => (
  <Icon viewBox="0 0 16 16" fill="none" color="green.600" {...props}>
    <path
      d="M2.76495 8.74786L5.75641 12.1667L13.235 3.83333"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </Icon>
);
