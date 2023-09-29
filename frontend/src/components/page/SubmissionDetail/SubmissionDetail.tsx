import { JudgeStatus as PbJudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { JudgeStatus } from "@/src/model/judge";
import { useRouterSubmissionId } from "@/src/usecases/contest";
import { useGetSubmissionDetail } from "@/src/usecases/judge";
import {
  Card,
  CardBody,
  CardHeader,
  Container,
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
import { JudgeStatusBadge } from "../../model/judge/JudgeStatusBadge";
import { SubmissionInfo } from "../../model/judge/SubmissionInfo";
import { Link } from "../../ui/Link";

const Editor = dynamic(() => import("@/src/components/ui/Editor").then(mod => mod.Editor), {
  loading: () => <p>読み込み中です</p>,
  ssr: false,
});

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
        エラーが発生しました <Link href="/">トップページへ戻る</Link>
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
            <SubmissionInfo
              submissionDetail={submissionDetail}
              minW="750px"
              fontSize="sm"
            />

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
