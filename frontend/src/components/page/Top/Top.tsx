import { useListContests } from "@/src/usecases/contest";
import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Grid,
  GridItem,
  Heading,
  HStack,
  List,
  ListItem,
  Text,
} from "@chakra-ui/react";
import { Link } from "../../ui/Link";

const ContestBanner = () => {
  const { contests, isLoading, error } = useListContests();

  if (isLoading) {
    return <>読み込み中</>;
  }

  if (error) {
    if (error.code === 5) {
      return <>コンテスト情報はありません</>;
    }
    return <>コンテスト情報の取得に失敗しました</>;
  }

  return (
    <>
      <List>
        {contests!.map((c, i) => (
          <ListItem key={i}>
            <Link href={`/contests/${c.slug}`}>{c.name}</Link>
          </ListItem>
        ))}
      </List>
      <Box pt={16}>
        <Link fontWeight="bold" href="/contests">コンテスト一覧</Link>
      </Box>
    </>
  );
};

export const Top = () => {
  return (
    <Grid
      minH="100%"
      gap={4}
      pb={4}
      templateAreas={`"hero hero hero"
                      "contests tasks info"`}
      templateRows="350px 1fr"
      templateColumns="repeat(3, 1fr)"
    >
      <GridItem area="hero" bg="#212121">
        <HStack gap={16} h="100%" justifyContent="center">
          {/** output: "export" にしていると next/image による image optimization が使えない */}
          <img src="/szpp-judge-logo.png" alt="logo image" width={293} height={293} />
          <Text fontSize="lg" lineHeight={7} color="white" whiteSpace="pre-wrap">
            {`SZPP Judge は静岡大学プログラミングサークル SZPP (すずっぷ) が運営する、
競技プログラミング用オンラインジャッジサイトです。

ビジュアルプログラミング言語 Scratch での提出にも対応しており、
市民向けプログラミング講座などで活用しています。`}
          </Text>
        </HStack>
      </GridItem>

      <GridItem area="contests" display="flex" pl={4}>
        <Card borderRadius={16} flex={1}>
          <CardHeader textAlign="center">
            <Heading size="md">最新のコンテスト</Heading>
          </CardHeader>
          <CardBody>
            <ContestBanner />
          </CardBody>
        </Card>
      </GridItem>

      <GridItem area="tasks" display="flex">
        <Card borderRadius={16} flex={1}>
          <CardHeader textAlign="center">
            <Heading size="md">最新の問題</Heading>
          </CardHeader>
          <CardBody>
            <Text>準備中</Text>
          </CardBody>
        </Card>
      </GridItem>

      <GridItem area="info" display="flex" pr={4}>
        <Card borderRadius={16} flex={1}>
          <CardHeader textAlign="center">
            <Heading size="md">お知らせ</Heading>
          </CardHeader>
          <CardBody>
            <Text>準備中</Text>
          </CardBody>
        </Card>
      </GridItem>
    </Grid>
  );
};
