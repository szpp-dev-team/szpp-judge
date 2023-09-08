import { JudgeStatusBadge } from "@/src/components/model/judge/JudgeStatusBadge";
import { DifficultyBadge } from "@/src/components/model/task/DifficultyBadge";
import { TaskWithMySubmissionSummary } from "@/src/model/task";
import { QuestionIcon } from "@chakra-ui/icons";
import {
  Box,
  Card,
  CardBody,
  CardHeader,
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
import { useRouter } from "next/router";
import { FC } from "react";
import { Link } from "../../ui/Link";

export type TaskCollectionProps = {
  tasks: TaskWithMySubmissionSummary[];
};

export const TaskCollection: FC<TaskCollectionProps> = ({ tasks }) => {
  const { query } = useRouter();
  const { contest_slug } = query;

  return (
    <Box px={16} h="inherit">
      <Card px={3} py={4} h="100%">
        <CardHeader>
          <Heading as="h1">問題一覧 / 得点状況</Heading>
        </CardHeader>
        <CardBody display="flex" flexDir="column" gap={4}>
          <Flex flexDir="column" flex={0}>
            <Text fontSize="md">現在の合計得点： 400 点</Text>
            <Flex alignItems="center" gap={2}>
              <Text>ペナルティ： +5 分</Text>
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
                  <Th textAlign="center">自分の結果</Th>
                  <Th textAlign="center">得点</Th>
                </Tr>
              </Thead>
              <Tbody>
                {tasks.map(t => (
                  <Tr key={t.id}>
                    <Td textAlign="center">
                      <Link href={`/contests/${contest_slug}/tasks/${t.slug}`}>{t.id}</Link>
                    </Td>
                    <Td textAlign="center">
                      <DifficultyBadge dif={t.difficulty} />
                    </Td>
                    <Td textAlign="center">{t.haiten}</Td>
                    <Td textAlign="left" whiteSpace="normal" minW={56}>
                      <Link href={`/contests/${contest_slug}/tasks/${t.slug}`}>{t.title}</Link>
                    </Td>
                    <Td textAlign="center">
                      {t?.status ? <JudgeStatusBadge status={t.status} progress={t?.progress} /> : "-"}
                    </Td>
                    <Td textAlign="center">{t?.score ?? "-"}</Td>
                  </Tr>
                ))}
              </Tbody>
            </Table>
          </TableContainer>
        </CardBody>
      </Card>
    </Box>
  );
};