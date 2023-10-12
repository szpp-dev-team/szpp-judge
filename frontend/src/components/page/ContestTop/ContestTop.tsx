import { useGetContest, useRegisterMe, useRouterContestSlug } from "@/src/usecases/contest";
import { Box, Button, Card, CardBody, CardHeader, Heading, Text, useToast } from "@chakra-ui/react";
import dynamic from "next/dynamic";
import { useCallback, useState } from "react";
import { useTimer } from "react-timer-hook";
import contestTopStyle from "./ContestTop.module.scss";

const MarkdownView = dynamic(() => import("@/src/components/ui/MarkdownView").then(mod => mod.MarkdownView), {
  ssr: false,
});

const TimerComponent = (props: { expiryTimestamp: Date; timerHeader: string; onExpire?: () => void }) => {
  const timer = useTimer({ expiryTimestamp: props.expiryTimestamp, onExpire: props.onExpire });
  // leading zero を追加する処理
  const seq = [
    timer.days.toLocaleString(),
    timer.hours.toLocaleString(),
    timer.minutes.toLocaleString(),
    timer.seconds.toLocaleString(),
  ];
  for (let i = 0; i < 4; i++) {
    if (seq[i]!.length == 1) seq[i] = "0" + seq[i];
  }
  return (
    <span>
      {props.timerHeader + seq[0] + ":" + seq[1] + ":" + seq[2] + ":" + seq[3]}
    </span>
  );
};

export const ContestTop = () => {
  const [key, setKey] = useState(1);
  const slug = useRouterContestSlug();
  const { contest } = useGetContest({ slug });

  // コンテスト開始まで or 終了までのタイマー
  let expiryTimestamp: Date;
  const now = new Date();
  let isUpcomingOrRunning = false;
  let timerHeader = "";
  if (contest?.startAt && now < contest.startAt.toDate()) {
    expiryTimestamp = contest.startAt.toDate();
    isUpcomingOrRunning = true;
    timerHeader = "開始まで ";
  } else if (contest?.endAt && now <= contest.endAt.toDate()) {
    expiryTimestamp = contest.endAt.toDate();
    isUpcomingOrRunning = true;
    timerHeader = "終了まで ";
  } else {
    expiryTimestamp = now;
  }

  const { isLoading, mutate } = useRegisterMe();
  const toast = useToast({ position: "bottom" });
  const handleClick = () => {
    console.info("button clicked");
    mutate(
      { contestSlug: slug },
      {
        onSuccess: () => {
          toast({ title: "参加登録しました", status: "success" });
        },
      },
    );
  };

  // HACK: タイマーの終了に合わせて isUpcomingOrRunning や timerHeader などを更新したい
  // 無理やりページごと React に re-render させるために key を書き換えている
  // もうちょいキレイなやり方はありそう…
  // tanstack query のおかげで getContest でリクエストが何度も飛ぶことはない
  const forceRerender = useCallback(() => {
    setKey(prev => prev + 1);
  }, []);

  return (
    <Box key={key} px={16} h="100%">
      <Card px={3} py={4} h="100%" borderRadius={0} color="cyan.900">
        <CardHeader textAlign="center">
          <Heading as="h1" mb={4}>{contest?.name}</Heading>
          <Text>
            開催期間: {contest?.startAt?.toDate().toLocaleString()} - {contest?.endAt?.toDate().toLocaleString()}
          </Text>
          {contest && isUpcomingOrRunning
            ? (
              <Box>
                <Button colorScheme="teal" onClick={handleClick} isLoading={isLoading} margin={4}>
                  参加登録
                </Button>
                <Text>
                  <TimerComponent
                    expiryTimestamp={expiryTimestamp}
                    timerHeader={timerHeader}
                    onExpire={forceRerender}
                  />
                </Text>
              </Box>
            )
            : <></>}
        </CardHeader>
        <CardBody>
          <MarkdownView markdown={contest?.description} className={contestTopStyle.markdownWrapper} />
        </CardBody>
      </Card>
    </Box>
  );
};
