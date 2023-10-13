import { Spinner, Text } from "@chakra-ui/react";
import { CenteringFullSizeContainer } from "./CenteringFullSizeContainer";

export const LoadingPane = () => {
  return (
    <CenteringFullSizeContainer flexDir="column" gap={3}>
      <Spinner
        thickness="6px"
        speed="0.7s"
        emptyColor="gray.200"
        color="teal.500"
        size="xl"
      />
      <Text>読み込み中...</Text>
    </CenteringFullSizeContainer>
  );
};
