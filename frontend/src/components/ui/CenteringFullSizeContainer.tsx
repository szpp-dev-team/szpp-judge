import { Center, CenterProps } from "@chakra-ui/react";

type CenteringFullSizeContainerProps = CenterProps;

export const CenteringFullSizeContainer = (props: CenteringFullSizeContainerProps) => {
  return <Center w="full" h="full" textAlign="center" p={2} {...props} />;
};
