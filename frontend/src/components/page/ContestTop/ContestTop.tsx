import { useGetContest, useRouterContestSlug } from "@/src/usecases/contest";
import {
  Box,
  Card,
  CardHeader,
  Heading,
  CardBody
} from "@chakra-ui/react";

export const ContestTop = () => {
  const slug = useRouterContestSlug();
  const { contest } = useGetContest({slug});
  return (
    <Box px={16} h="100%">
      <Card px={3} py={4} h="100%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">{contest?.name}</Heading>
        </CardHeader>
        <CardBody>
          <p>
            開催期間: {contest?.startAt?.toDate().toLocaleString()} - {contest?.endAt?.toDate().toLocaleString()}
          </p>
          <p>{contest?.description}</p>
        </CardBody>
      </Card>
    </Box>
  );
};
