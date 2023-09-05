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
  Thead,
  Tr,
} from "@chakra-ui/react";
import { FC } from "react";

export type TaskCollectionProps = {
  tasks: TaskWithMySubmissionSummary[];
};

export const TaskCollection: FC<TaskCollectionProps> = ({ tasks }) => {
  return (
    <Box px={16}>
      <Card px={3} py={4}>
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
            <Table variant="unstyled">
              <Thead>
                <Tr>
                  <Td>#</Td>
                  <Td>難易度</Td>
                  <Td>配点</Td>
                  <Td>問題名</Td>
                  <Td>自分の結果</Td>
                  <Td>得点</Td>
                </Tr>
              </Thead>
              <Tbody>
                {tasks.map(t => (
                  <Tr key={t.id}>
                    <Td>{t.id}</Td>
                    <Td>{t.difficulty}</Td>
                    <Td>{t.haiten}</Td>
                    <Td>{t.title}</Td>
                    <Td>{t?.status ?? "-"}</Td>
                    <Td>{t?.score ?? "-"}</Td>
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
