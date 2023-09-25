import { Link } from "@/src/components/ui/Link";
import { useListContests } from "@/src/usecases/contest";
import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Heading,
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  Td
} from "@chakra-ui/react";
import { Contest } from "@/src/gen/proto/backend/v1/contest_resources_pb";

const ContestListWithFilter = (header: string, f: (c: Contest) => boolean) => {
  const { contests } = useListContests({});
  return (
    <Card px={3} py={4} borderRadius={0}>
      <CardHeader>
        <Heading as="h2">{header}</Heading>
      </CardHeader>
      <CardBody>
        <TableContainer>
          <Table variant="bordered-narrow">
            <Thead>
              <Tr>
                <Th textAlign="center" width={100}>開始</Th>
                <Th textAlign="center" width={100}>終了</Th>
                <Th textAlign="center">コンテスト名</Th>
                <Th textAlign="center">問題数</Th>
              </Tr>
            </Thead>
            <Tbody>
              {contests?.filter(f).map((c) => (
                <Tr key={c.slug}>
                  <Td textAlign="center">{c.startAt?.toDate().toLocaleString()}</Td>
                  <Td textAlign="center">{c.endAt?.toDate().toLocaleString()}</Td>
                  <Td textAlign="left">
                    <Link href={`/contests/${c.slug}`}>{c.name}</Link>
                  </Td>
                  <Td textAlign="center">{c.taskIds.length}</Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </TableContainer>
      </CardBody>
    </Card>
  );
};

export const ContestList = () => {
  const now = new Date();
  return (
    <Box px={16} h="100%">
      <Card px={3} py={4} h="100%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">コンテスト一覧</Heading>
        </CardHeader>
        <CardBody>
          {ContestListWithFilter("開催中", (c) => {
            return c.startAt != undefined && c.startAt?.toDate() <= now
                   && c.endAt != undefined && now <= c.endAt?.toDate();
          })}
          {ContestListWithFilter("開催予定", (c) => {
            return c.startAt != undefined && now < c.startAt?.toDate();
          })}
          {ContestListWithFilter("終了済み", (c) => {
            return c.endAt != undefined && c.endAt?.toDate() < now;
          })}
        </CardBody>
      </Card>
    </Box>
  );
};
