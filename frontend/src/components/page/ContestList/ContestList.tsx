import { Link } from "@/src/components/ui/Link";
import { useListContests } from "@/src/usecases/contest";
import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Heading
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
          <div>
            <strong>動作確認用</strong>
            <ul>
              <li>
                <Link href="/contests/not_exist/tasks">
                  Go to <code>note_exist/tasks</code>
                </Link>
              </li>
              {contests?.map((c) => (
                <li key={c.id}>
                  <Link href={`/contests/${c.slug}`}>Go to {c.name}</Link>
                </li>
              ))}
            </ul>
          </div>
        </CardBody>
      </Card>
    </Box>
  );
};
