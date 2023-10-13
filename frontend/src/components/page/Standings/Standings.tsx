import { Link } from "@/src/components/ui/Link";
import { useGetContest, useRouterContestSlug, useStandings, useTaskSeqCodes } from "@/src/usecases/contest";
import { Duration } from "@/src/util/time";
import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Container,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  VStack,
} from "@chakra-ui/react";

type TaskScoreProps = {
  /** タスク(への提出)の得点 */
  score?: number;
  penaltyCount?: number;
  untilAc?: number | bigint;
  /**
   * AC した提出へのリンク
   * (渡されると表示しちゃうのでコンテスト開催中は渡さないようにする)
   */
  href?: string;
};

type TotalScoreProps = {
  /** コンテストの得点 */
  score?: number;
  penaltyCount?: number;
  /** ペナルティを加味した回答にかかった時間. protobuf 的には Duration だが諸事情によりタイムスタンプっぽい名前になっている */
  latestAcAt?: number | bigint; // TODO: proto とともに名前を直す
};

/**
 * 順位表のテーブルのセルに入れるための得点コンポーネント(総得点)
 */
const TotalScore = ({ score, penaltyCount, latestAcAt }: TotalScoreProps) => {
  /** 1つ以上の AC を提出しているかどうか(提出があっても AC でなければだめ) */
  const hasAc = typeof score === "number" && score > 0;
  if (!hasAc) {
    return <Text as="span" color="gray.400">-</Text>;
  }

  return (
    <VStack gap={0}>
      <Box as="p">
        <Text as="span" color="blue.500" fontWeight="bold">{score}</Text>
        {penaltyCount
          ? <Text as="span" fontWeight="bold" color="pink.400">({penaltyCount})</Text>
          : null}
      </Box>
      {latestAcAt ? <Text color="gray.400">{Duration.fromNumber(latestAcAt).fmtHMS()}</Text> : null}
    </VStack>
  );
};

/**
 * 順位表のテーブルのセルに入れるための得点コンポーネント(タスク毎)
 */
const TaskScore = ({ score, penaltyCount, untilAc, href }: TaskScoreProps) => {
  /** 1つ以上の AC を提出しているかどうか(提出があっても AC でなければだめ) */
  const hasAc = typeof score === "number" && score > 0;
  if (!hasAc) {
    return <Text as="span" color="gray.400">-</Text>;
  }

  const sharedProps = {
    color: "green.500",
    fontWeight: "bold",
  };

  return (
    <VStack gap={0}>
      <Box as="p">
        {href ? <Link href={href} {...sharedProps}>{score}</Link> : <Text as="span" {...sharedProps}>{score}</Text>}
        {penaltyCount ? <Text as="span" fontWeight="bold" color="pink.400">({penaltyCount})</Text> : null}
      </Box>
      {untilAc ? <Text color="gray.400">{Duration.fromNumber(untilAc).fmtHMS()}</Text> : null}
    </VStack>
  );
};

export const Standings = () => {
  const contestSlug = useRouterContestSlug();

  const { standingsList, error, isLoading } = useStandings({ contestSlug });
  const tasks = standingsList?.length ? standingsList[0]!.taskDetailList : [];

  const seqCodes = useTaskSeqCodes(tasks.length);

  const { contest, error: cError, isLoading: cIsLoading } = useGetContest({ slug: contestSlug });
  // REVIEW: コンテスト終了間際にページを開いてコンテスト終了後リロードせずに true に変わってほしいができてるか？
  const isSubmissionVisible = ((c: typeof contest) => {
    if (!c) return undefined;
    const now = new Date();
    const { startAt, endAt } = c;
    return startAt && endAt && startAt.toDate() < now && endAt.toDate() < now;
  })(contest);

  // HACK: 本来は useStandings だけ面倒見ればいいのに useGetContest の方もハンドリングするの微妙
  if (isLoading || cIsLoading) {
    return <>読み込み中です...</>;
  }

  if (error || cError) {
    return (
      <>
        エラーが発生しました <Link href="/">トップページへ戻る</Link>
      </>
    );
  }

  return (
    <Container maxW="container.lg">
      <Card px={3} py={4} h="100%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">順位表</Heading>
        </CardHeader>

        <CardBody>
          {standingsList && standingsList.length
            ? (
              <TableContainer>
                <Table variant="bordered-narrow">
                  <Thead>
                    <Tr>
                      <Th key="th-rank" textAlign="center">順位</Th>
                      <Th key="th-user" textAlign="center">ユーザ</Th>
                      <Th key="th-totalscore" textAlign="center">得点</Th>
                      {tasks.map((_task, i) => (
                        <Th key={i} textAlign="center">
                          <Link href={`/contests/${contestSlug}/tasks/${i + 1}`}>
                            {seqCodes[i]}
                          </Link>
                        </Th>
                      ))}
                    </Tr>
                  </Thead>

                  <Tbody>
                    {standingsList.map((record, i) => (
                      <Tr key={i}>
                        <Td key={`${i}-rank`} textAlign="center">{record.rank}</Td>
                        <Td key={`${i}-user`} textAlign="center">{record.username}</Td>
                        <Td key={`${i}-totalscore`} textAlign="center">
                          <TotalScore
                            score={record.totalScore}
                            penaltyCount={record.totalPenaltyCount}
                            latestAcAt={record.latestAcAt
                              ? record.latestAcAt.seconds * BigInt(Duration.SECOND)
                              : undefined}
                          />
                        </Td>
                        {record.taskDetailList.map((task, j) => (
                          <Td key={`${i}-task${j}`} textAlign="center">
                            <TaskScore
                              score={task.score}
                              penaltyCount={task.penaltyCount}
                              untilAc={task.untilAc ? task.untilAc.seconds * BigInt(Duration.SECOND) : undefined}
                              href={typeof isSubmissionVisible === "boolean" && isSubmissionVisible
                                  && typeof task.acSubmitId !== "undefined"
                                ? `/contests/${contestSlug}/submissions/${task.acSubmitId}`
                                : undefined}
                            />
                          </Td>
                        ))}
                      </Tr>
                    ))}
                  </Tbody>
                </Table>
              </TableContainer>
            )
            : <Text>順位表はまだありません</Text>}
        </CardBody>
      </Card>
    </Container>
  );
};
