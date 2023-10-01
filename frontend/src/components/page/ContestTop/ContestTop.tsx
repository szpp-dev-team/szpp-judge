import { useGetContest, useRouterContestSlug } from "@/src/usecases/contest";
import {
  Box,
  Card,
  CardHeader,
  Heading,
  CardBody,
  Button,
} from "@chakra-ui/react";
import { useAccessTokenClaimValue } from "@/src/globalStates/credential";
import { useTimer } from "react-timer-hook";

export const ContestTop = () => {
  const slug = useRouterContestSlug();
  const { contest } = useGetContest({slug});

  // コンテスト開始まで or 終了までのタイマー
  // 現状うまく動かない　現在時刻基準で 10 秒後にセットとかすると問題なく動くが、contest::endAt とかを突っ込むと一生 {0, 0, 0, 0} が返る
  // コンテスト情報がダミーなのが悪さをしている？
  let expiryTimestamp: Date;
  const now = new Date();
  let isUpcomingOrRunning = false;
  let timerHeader = ""
  if (contest?.startAt != undefined && now < contest.startAt.toDate()) {
    expiryTimestamp = contest.startAt.toDate();
    isUpcomingOrRunning = true;
    timerHeader = "コンテスト開始まで ";
  } else if (contest?.endAt != undefined && now <= contest.endAt.toDate()) {
    expiryTimestamp = contest.endAt.toDate();
    isUpcomingOrRunning = true;
    timerHeader = "コンテスト終了まで ";
  } else {
    expiryTimestamp = now;
  }
  //expiryTimestamp = now;
  //expiryTimestamp.setSeconds(expiryTimestamp.getSeconds() + 10);
  const timer = useTimer({expiryTimestamp});

  // leading zero を追加する処理
  let seq = [
    timer.days.toLocaleString(),
    timer.hours.toLocaleString(),
    timer.minutes.toLocaleString(),
    timer.seconds.toLocaleString()
  ];
  for (let i = 0; i < 4; i++) {
    if (seq[i].length == 1) seq[i] = "0" + seq[i];
  }

  const user = useAccessTokenClaimValue();
  if (isUpcomingOrRunning && user != null) {
    // 参加登録しているかどうかを調べて生やすボタンの種類を決める
  }

  return (
    <Box px={16} h="100%">
      <Card px={3} py={4} h="100%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">{contest?.name}</Heading>
          {isUpcomingOrRunning
            ? <p>{timerHeader +  seq[0] + ":" + seq[1] + ":" + seq[2] + ":" + seq[3]}</p>
            : <></>
          }
        </CardHeader>
        <CardBody>
          <p>
            開催期間: {contest?.startAt?.toDate().toLocaleString()} - {contest?.endAt?.toDate().toLocaleString()}
          </p>
          <p>{contest?.description}</p>
          {isUpcomingOrRunning
            ? <Button colorScheme="teal">
                参加登録
              </Button>
            : <></>
          }
        </CardBody>
      </Card>
    </Box>
  );
};
