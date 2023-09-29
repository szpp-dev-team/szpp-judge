import { JudgeStatus as PbJudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { JudgeStatus } from "@/src/model/judge";
import { useRouterSubmissionId } from "@/src/usecases/contest";
import { useGetSubmissionDetail } from "@/src/usecases/judge";
import {
  Card,
  CardBody,
  CardHeader,
  Container,
  Grid,
  GridItem,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import dynamic from "next/dynamic";
import { ReactNode } from "react";
import { JudgeStatusBadge } from "../../model/judge/JudgeStatusBadge";
import { Link } from "../../ui/Link";

const Editor = dynamic(() => import("@/src/components/ui/Editor").then(mod => mod.Editor), {
  loading: () => <p>読み込み中です</p>,
  ssr: false,
});

type PairProps = {
  data: [/** key */ string, /** value */ ReactNode];
  //  gridRow?: GridProps["gridRow"],
};

const Pair = ({ data }: PairProps): ReactNode => {
  return (
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
      <Container maxW="container.lg">
        <Card px={3} py={4} h="100%" borderRadius={0}>
          <CardHeader>
            <Heading as="h1">提出 #{idStr}</Heading>
          </CardHeader>
          <CardBody display="flex" flexDir="column" gap={4}>
            <Heading as="h2" size={["lg", null, "lg"]}>提出情報</Heading>
            <Grid
              as="dl"
              textAlign="center"
              templateColumns="3fr 5fr"
              borderTop="1px"
              borderLeft="1px"
              borderColor="gray.300"
              fontSize="sm"
              minW="750px"
            >
              {/* sv-SE ローケルでのフォーマットは `YYYY-mm-dd HH:MM:SS`*/}
              <Pair data={["提出日時", submissionDetail?.submittedAt?.toDate().toLocaleString("sv-SE")]} />
              <Pair data={["問題", <Link href="#">{submissionDetail?.taskTitle}</Link>]} />
              <Pair
                data={[
                  "ユーザ",
                  <Link href={`/users/${submissionDetail?.username}`}>{submissionDetail?.username}</Link>,
                ]}
              />
              <Pair data={["言語", `${submissionDetail?.execTimeMs} ms`]} />
              <Pair data={["得点", submissionDetail?.score]} />
              <Pair
                data={[
                  "結果",
                  <JudgeStatusBadge
                    status={PbJudgeStatus[submissionDetail?.status ?? 0].toString() as JudgeStatus}
                  />,
                ]}
              />
              <Pair data={["実行時間", `${submissionDetail?.execTimeMs} ms`]} />
              <Pair data={["メモリ", `${submissionDetail?.execMemoryKib} KB`]} />
            </Grid>

            <Heading as="h2" size={["md", null, "lg"]}>ジャッジ結果</Heading>
            <TableContainer>
              <Table variant="bordered-narrow">
                <Thead>
                  <Tr>
                    <Th textAlign="center">ケース名</Th>
                    <Th textAlign="center">結果</Th>
                    <Th textAlign="center">実行時間</Th>
                    <Th textAlign="center">メモリ</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {submissionDetail?.testcaseResults.map((t, i) => (
                    <Tr key={i}>
                      <Td textAlign="center">{t.testcaseName}</Td>
                      <Td textAlign="center">
                        <JudgeStatusBadge
                          status={PbJudgeStatus[t.judgeStatus].toString() as JudgeStatus}
                        />
                      </Td>
                      <Td textAlign="center">{t.execTimeMs} ms</Td>
                      <Td textAlign="center">{t.execMemoryKib} KB</Td>
                    </Tr>
                  ))}
                </Tbody>
              </Table>
            </TableContainer>
            <Heading as="h2" size={["md", null, "lg"]}>ソースコード</Heading>
            <Editor doc={submissionDetail?.sourceCode} readonly={true} />
          </CardBody>
        </Card>
      </Container>
    </>
  );
};
