import { useGetContest, useRouterContestSlug } from "@/src/usecases/contest";
import {
  Box,
  Card,
  CardHeader,
  Heading,
  CardBody,
  Button,
} from "@chakra-ui/react";
//import { useTimer } from "react-timer-hook";

export const ContestTop = () => {
  const slug = useRouterContestSlug();
  const { contest } = useGetContest({slug});

  /*let expiryTimestamp: Date;
  const date = new Date();
  if (contest?.startAt != undefined && date < contest.startAt.toDate()) {
    expiryTimestamp = contest.startAt.toDate();
  } else if (contest?.endAt != undefined && date <= contest.endAt.toDate()) {
    expiryTimestamp = contest.endAt.toDate();
  } else {
    expiryTimestamp = date;
  }
  expiryTimestamp = date;
  expiryTimestamp.setSeconds(expiryTimestamp.getSeconds() + 10);

  const timer = useTimer({expiryTimestamp});

  let seq = [
    timer.days.toLocaleString(),
    timer.hours.toLocaleString(),
    timer.minutes.toLocaleString(),
    timer.seconds.toLocaleString()
  ];
  for (let i = 0; i < 4; i++) {
    if (seq[i].length == 1) seq[i] = "0" + seq[i];
  }*/

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
          <Button colorScheme="teal">
            参加登録
          </Button>
        </CardBody>
      </Card>
    </Box>
  );
};
