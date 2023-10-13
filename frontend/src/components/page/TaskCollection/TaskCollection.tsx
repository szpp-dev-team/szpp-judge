import { DifficultyBadge } from "@/src/components/model/task/DifficultyBadge";
import { taskConverter } from "@/src/model/converter/taskConverter";
import { ScoreStatus } from "@/src/model/task";
import {
  useGetContest,
  useGetMySubmissionStatuses,
  useIsContestStarted,
  useListContestTasks,
  useRouterContestSlug,
} from "@/src/usecases/contest";
import { QuestionIcon } from "@chakra-ui/icons";
import {
  Card,
  CardBody,
  CardHeader,
  Container,
  Flex,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { useMemo } from "react";
import { Link } from "../../ui/Link";

const getBgColorFromScoreStatus = (s: ScoreStatus): string =>
  s == null ? "initial" : s === "perfect" ? "#c3e6cb" : "#ffeeba";

export const TaskCollection = () => {
  const contestSlug = useRouterContestSlug();
  const { contest, error: errorC } = useGetContest({ slug: contestSlug });

  const isContestStarted = useIsContestStarted({ contestSlug, now: new Date() });
  const { tasks, error: errorT } = useListContestTasks({ contestSlug }, { isContestStarted });

  const { submissionStatuses, error: errorS } = useGetMySubmissionStatuses({ contestSlug }, { isContestStarted });

  const tasksWithScore = useMemo(() => {
    if (tasks == null || submissionStatuses == null) return [];
    return tasks.map((t) => {
      // submissionStatuses には提出したことのある問題の結果しか含まれない
      const ss = submissionStatuses.find(s => s.taskId === t.id)?.score;
      return {
        ...t,
        submissionScore: ss,
        scoreStatus: ScoreStatus.fromScore(t.score, ss),
      };
    });
  }, [tasks, submissionStatuses]);

  if (errorC || errorT || errorS) {
    return <>エラーが発生しました</>;
  }

  const totalScore = submissionStatuses?.reduce((prev, s) => prev + (s.score ?? 0), 0) || 0;
  const penaltyMinutes = typeof contest?.penaltySeconds === "number" ? Math.round(contest.penaltySeconds / 60) : "?";

  return (
    <Container h="100%" maxW="container.lg">
      <Card px={3} py={4} h="100%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">問題一覧 / 得点状況</Heading>
        </CardHeader>
        <CardBody display="flex" flexDir="column" gap={4}>
          <Flex flexDir="column" flex={0}>
            <Text fontSize="md">現在の合計得点： {totalScore} 点</Text>
            <Flex alignItems="center" gap={2}>
              <Text>ペナルティ： +{penaltyMinutes} 分</Text>
              <QuestionIcon color="cyan.600" />
            </Flex>
          </Flex>
          <TableContainer>
            <Table variant="bordered-narrow">
              <Thead>
                <Tr>
                  <Th textAlign="center">#</Th>
                  <Th textAlign="center">難易度</Th>
                  <Th textAlign="center">配点</Th>
                  <Th textAlign="left">問題名</Th>
                  <Th textAlign="center">得点</Th>
                </Tr>
              </Thead>
              <Tbody>
                {tasksWithScore.map((t, seq) => (
                  <Tr key={t.id} backgroundColor={getBgColorFromScoreStatus(t.scoreStatus)}>
                    <Td textAlign="center">
                      <Link href={`/contests/${contestSlug}/tasks/${seq + 1}`}>{seq + 1}</Link>
                    </Td>
                    <Td textAlign="center">
                      <DifficultyBadge dif={taskConverter.difficulty.from(t.difficulty)} />
                    </Td>
                    <Td textAlign="center">{t.score}</Td>
                    <Td textAlign="left" whiteSpace="normal" minW={56}>
                      <Link href={`/contests/${contestSlug}/tasks/${seq + 1}`}>{t.title}</Link>
                    </Td>
                    <Td textAlign="center">{t.submissionScore}{/* undefined だと空欄になる */}</Td>
                  </Tr>
                ))}
              </Tbody>
            </Table>
          </TableContainer>
        </CardBody>
      </Card>
    </Container>
  );
};
