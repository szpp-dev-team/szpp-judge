import type { ScoreStatus } from "@/src/model/task";
import { Icon, type IconProps } from "@chakra-ui/react";
import { RoundedCheckIcon } from "../../icon/RoundedCheckIcon";
import { RoundedCrossIcon } from "../../icon/RoundedCrossIcon";
import { RoundedTriangleIcon } from "../../icon/RoundedTriangleIcon";

export type ScoreStatusIconProps = {
  status: ScoreStatus;
} & IconProps;

export const ScoreStatusIcon = ({ status, ...props }: ScoreStatusIconProps) => {
  if (status == null) {
    return (
      <Icon viewBox="0 0 16 16" fill="none" {...props}>
        <path />
      </Icon>
    );
  }
  switch (status) {
    case "perfect":
      return <RoundedCheckIcon {...props} />;
    case "partial":
      return <RoundedTriangleIcon {...props} />;
    case "zero":
      return <RoundedCrossIcon {...props} />;
    default: {
      const exhaustiveCheck: never = status;
      throw new Error(`Invalid scoreStatus: ${exhaustiveCheck}`);
    }
  }
};
