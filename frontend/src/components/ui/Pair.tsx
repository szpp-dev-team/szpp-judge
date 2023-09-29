import { GridItem } from "@chakra-ui/react";
import type { ReactNode } from "react";

type PairProps = {
  data: [/** key */ string, /** value */ ReactNode];
};

export const Pair = ({ data }: PairProps) => (
  <>
    <GridItem
      as="dt"
      display="flex"
      py={2}
      alignItems="center"
      justifyContent="center"
      borderBottom="1px"
      borderRight="1px"
      borderColor="gray.300"
      bg="gray.50"
      fontWeight="bold"
    >
      {data[0]}
    </GridItem>
    <GridItem
      as="dd"
      display="flex"
      py={2}
      alignItems="center"
      justifyContent="center"
      borderBottom="1px"
      borderRight="1px"
      borderColor="gray.300"
    >
      {data[1]}
    </GridItem>
  </>
);
