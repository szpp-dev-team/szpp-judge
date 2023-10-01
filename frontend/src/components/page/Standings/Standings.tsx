import { Link } from "@/src/components/ui/Link";
import { calcNthTaskSeq, useIsContestClosed, useRouterContestSlug, useStandings } from "@/src/usecases/contest";
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
import { useMemo } from "react";

type TaskScoreProps = {
  /** タスク(への提出)の得点 */
  score?: number;
  penaltyCount?: number;
  untilAc?: number | bigint;
  acSubmitId?: number;
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
  untilAc?: number | bigint;
};

/**
 * 順位表のテーブルのセルに入れるための得点コンポーネント(総得点)
 */
const TotalScore = ({ score, penaltyCount, untilAc, ...rest }: TotalScoreProps) => {
  /** 1つ以上の AC を提出しているかどうか */
  const hasAc = typeof score === "number" && score > 0;
  if (!hasAc) {
    return <Text as="span" color="gray.400">-</Text>;
  }

  return (
    <VStack gap={0} {...rest}>
      <Box as="p">
        <Text as="span" color="blue.500" fontWeight="bold">{score}</Text>
        {penaltyCount
          ? <Text as="span" fontWeight="bold" color="pink.400">({penaltyCount})</Text>
          : null}
      </Box>
      {untilAc ? <Text color="gray.400">{Duration.fromNumber(untilAc).fmtHMS()}</Text> : null}
    </VStack>
  );
};

/**
 * 順位表のテーブルのセルに入れるための得点コンポーネント(タスク毎)
 */
const TaskScore = ({ score, penaltyCount, untilAc, href, ...rest }: TaskScoreProps) => {
  const hasAc = typeof score === "number" && score > 0;
  if (!hasAc) {
    return <Text as="span" color="gray.400">-</Text>;
  }

  const sharedProps = {
    color: "green.500",
    fontWeight: "bold",
  };

  return (
    <VStack gap={0} {...rest}>
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
  let contestId = 1; // TODO: contestSlug から standings を取れるようにする
  try {
    contestId = Number.parseInt(contestSlug.slice(-1), 10); // "sbc005" -> Number(5)
  } catch { /* empty */ }

  const { standingsList, error, isLoading } = useStandings({ contestId });
  const tasks = useMemo(() => {
    if (standingsList == null || standingsList.length === 0) {
      return [];
    } else {
      return standingsList[0].taskDetailList;
    }
  }, [standingsList]);
  const isContestClosed = useIsContestClosed({ slug: contestSlug });

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
    <Container maxW="container.lg">
      <Card px={3} py={4} h="10%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">順位表</Heading>
        </CardHeader>

        <CardBody>
          <TableContainer>
            <Table variant="bordered-narrow">
              <Thead>
                <Tr>
                  <Th key="th-rank" textAlign="center">順位</Th>
                  <Th key="th-user" textAlign="center">ユーザ</Th>
                  <Th key="th-totalscore" textAlign="center">得点</Th>
                  {tasks.map((_task, seq, allTask) => (
                    <Th key={seq} textAlign="center">
                      <Link href={`/contests/${contestSlug}/tasks/${seq}`}>{calcNthTaskSeq(seq, allTask.length)}</Link>
                    </Th>
                  ))}
                </Tr>
              </Thead>

              <Tbody>
                {standingsList?.map((record, i) => {
                  const taskDetailList = record.taskDetailList;
                  return (
                    <Tr key={i}>
                      <Td key={`${i}-rank`} textAlign="center">{record.rank}</Td>
                      <Td key={`${i}-user`} textAlign="center">{record.username}</Td>
                      <Td key={`${i}-totalscore`} textAlign="center">
                        <TotalScore
                          score={record.totalScore}
                          penaltyCount={record.totalPenaltyCount}
                          untilAc={1800000} // TODO: ペナルティを加味した untilAc 的な指標をバックエンドからもらう
                        />
                      </Td>
                      {taskDetailList.map((task, j) => (
                        <Td key={`${i}-task${j}`} textAlign="center">
                          <TaskScore
                            score={task.score}
                            penaltyCount={task.penaltyCount}
                            untilAc={task.untilAc ? task.untilAc.seconds * BigInt(Duration.SECOND) : undefined}
                            acSubmitId={task?.acSubmitId}
                            href={typeof isContestClosed === "boolean" && isContestClosed
                                && typeof task.acSubmitId !== "undefined"
                              ? `/contests/${contestSlug}/submissions/${task.acSubmitId}`
                              : undefined}
                          />
                        </Td>
                      ))}
                    </Tr>
                  );
                })}
              </Tbody>
            </Table>
          </TableContainer>
        </CardBody>
      </Card>
    </Container>
  );
};
