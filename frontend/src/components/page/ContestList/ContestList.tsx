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

export const ContestList = () => {
  const { contests } = useListContests({});
  return (
    <Box px={16} h="100%">
      <Card px={3} py={4} h="100%" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">コンテスト一覧</Heading>
        </CardHeader>
        <CardBody>
          <TableContainer>
            <Table variant="bordered-narrow">
              <Thead>
                <Tr>
                  <Th textAlign="center">開始日時</Th>
                  <Th textAlign="center">コンテスト名</Th>
                  <Th textAlign="center">問題数</Th>
                </Tr>
              </Thead>
              <Tbody>
                {contests?.map((c) => (
                  <Tr key={c.slug}>
                    <Td textAlign="center">{c.startAt?.toDate().toLocaleString()}</Td>
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
    </Box>
  );
};
