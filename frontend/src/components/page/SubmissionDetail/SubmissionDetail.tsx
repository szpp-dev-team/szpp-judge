import { JudgeStatus as PbJudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { JudgeStatus } from "@/src/model/judge";
import { useRouterSubmissionId } from "@/src/usecases/contest";
import { useGetSubmissionDetail } from "@/src/usecases/judge";
import { Box, Card, CardBody, CardHeader, Grid, GridItem, Heading } from "@chakra-ui/react";
import { ReactNode } from "react";
import { JudgeStatusBadge } from "../../model/judge/JudgeStatusBadge";
import { Link } from "../../ui/Link";

type PairProps = {
  data: [/** key */ string, /** value */ ReactNode];
  //  gridRow?: GridProps["gridRow"],
};

const Pair = ({ data }: PairProps): ReactNode => {
  return (
    <>
      <GridItem
        as="dt"
        fontSize="sm"
        display="flex"
        py={2}
        alignItems="center"
        justifyContent="center"
        borderWidth="thin"
        borderColor="gray.300"
        bg="gray.50"
        fontWeight="bold"
      >
        {data[0]}
      </GridItem>
      <GridItem
        as="dd"
        fontSize="sm"
        display="flex"
        py={2}
        alignItems="center"
        justifyContent="center"
        borderWidth="thin"
        borderColor="gray.300"
      >
        {data[1]}
      </GridItem>
    </>
  );
};

export const SubmissionDetail = () => {
  const idStr = useRouterSubmissionId();
  const id = Number.parseInt(idStr, 10); // NaN に注意
  const { submissionDetail, error, isLoading } = useGetSubmissionDetail({ id });

  if (isLoading) {
    return <>読み込み中です...</>;
  }

  if (error) {
    return (
      <>
        エラーが発生しました <a href="/">トップページへ戻る</a>
      </>
    );
  }

  return (
    <>
      <Box px={16} h="100%">
        <Card px={3} py={4} h="100%" borderRadius={0}>
          <CardHeader>
            <Heading as="h1">{`提出 #${idStr}`}</Heading>
          </CardHeader>
          <CardBody>
            <Grid as="dl" textAlign="center" templateColumns="3fr 5fr" borderWidth="thin" minW="750px">
              {/* sv-SE ローケルでのフォーマットは `YYYY-mm-dd HH:MM:SS`*/}
              <Pair data={["提出日時", submissionDetail?.submittedAt?.toDate().toLocaleString("sv-SE")]} />
              <Pair data={["問題", <Link href="#">{submissionDetail?.taskTitle}</Link>]} />
              <Pair data={["言語", `${submissionDetail?.execTimeMs} ms`]} />
              <Pair
                data={[
                  "ユーザ",
                  <Link href={`/users/${submissionDetail?.username}`}>{submissionDetail?.username}</Link>,
                ]}
              />
              <Pair data={["実行時間", `${submissionDetail?.execTimeMs} ms`]} />
              <Pair data={["メモリ", `${submissionDetail?.execMemoryKib} KB`]} />
              <Pair
                data={[
                  "結果",
                  <JudgeStatusBadge
                    status={PbJudgeStatus[submissionDetail?.status ?? 0].toString() as JudgeStatus}
                  />,
                ]}
              />
              <Pair data={["得点", submissionDetail?.score]} />
            </Grid>
          </CardBody>
        </Card>
      </Box>
    </>
  );
};
