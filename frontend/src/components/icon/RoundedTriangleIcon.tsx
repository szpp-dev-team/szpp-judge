import { Icon, IconProps } from "@chakra-ui/react";

export type RoundedTriangleIconProps = IconProps;

export const RoundedTriangleIcon = (props: RoundedTriangleIconProps) => (
  <Icon viewBox="0 0 16 16" fill="none" color="#E57106" {...props}>
    <path
      d="M8 3.33333L13.1962 12.3333L2.80385 12.3333L8 3.33333Z"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinejoin="round"
    />
  </Icon>
);
