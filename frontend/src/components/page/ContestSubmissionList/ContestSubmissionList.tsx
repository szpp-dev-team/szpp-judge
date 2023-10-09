import { LangID, langMetasBrief } from "@/src/gen/langs";
import { JudgeStatus as PbJudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { useAccessTokenClaimValue } from "@/src/globalStates/credential";
import { JudgeStatus } from "@/src/model/judge";
import { useGetContest, useRouterContestSlug } from "@/src/usecases/contest";
import { useGetJudgeProgress, useListSubmissions } from "@/src/usecases/judge";
import { fmtDatetime } from "@/src/util/time";
import {
  Card,
  CardBody,
  CardHeader,
  Container,
  Heading,
  Table,
  TableContainer,
  TableProps,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { useMemo } from "react";
import { JudgeStatusBadge } from "../../model/judge/JudgeStatusBadge";
import { Link } from "../../ui/Link";

/**
 * status の更新に追従するバッジ. 現在のステータスが WJ (ジャッジ待ち)
 * であることが明らかな submission についてのみ使用すること. すでに結果が
 * 確定している submission に使ってもリクエストの無駄なため.
 */
const ReactiveStatusCell = ({ submissionId }: { submissionId: number }) => {
  const { data, error, isLoading } = useGetJudgeProgress({ submissionId });
  if (isLoading) {
    return null;
  } else if (error) {
    console.error(error);
    return <>エラー</>;
  }

  const judgeStatus = typeof data.judgeProgress?.status === "number"
    ? PbJudgeStatus[data.judgeProgress.status] as JudgeStatus
    : "WJ";
  return <JudgeStatusBadge status={judgeStatus} />;
};

type SubmissionTableProps = {
  contestSlug: string;
  contestId: number; // 本当は contestSlug だけにしたい
  username?: string;
} & TableProps;

const SubmissionTable = ({ contestSlug, contestId, username, ...props }: SubmissionTableProps) => {
  // username は undefined の可能性あり
  const { data, isLoading, error } = useListSubmissions({ contestId, username });

  if (isLoading) return <>Now loading...</>;
  if (error) return <>Error occured</>;

  if (data.submissions.length === 0) {
    return (
      <Text>
        該当する提出がありません。左の<Link href={`/contests/${contestSlug}/tasks`}>
          問題
        </Link>から問題を探して提出してみてください。
      </Text>
    );
  }

  const judgeResult = (submissionId: number, fixedOrWJ: PbJudgeStatus | undefined) => {
    if (typeof fixedOrWJ === "undefined") {
      return null;
    }

    if (fixedOrWJ === PbJudgeStatus.WJ) {
      // waiting judge and request judge progress
      return <ReactiveStatusCell submissionId={submissionId} />;
    } else {
      // status is fixed
      const fixedStatus = PbJudgeStatus[fixedOrWJ] as JudgeStatus;
      return <JudgeStatusBadge status={fixedStatus} />;
    }
  };

  return (
    <TableContainer>
      <Table {...props} variant="bordered-narrow">
        <Thead>
          <Tr>
            <Th textAlign="center">提出日時</Th>
            <Th textAlign="center">問題</Th>
            <Th textAlign="center">ユーザ</Th>
            <Th textAlign="center">言語</Th>
            <Th textAlign="center">得点</Th>
            <Th textAlign="center">結果</Th>
            <Th textAlign="center">実行時間</Th>
            <Th textAlign="center">メモリ</Th>
            <Th textAlign="center">#</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.submissions.map((s, i) => (
            <Tr key={i}>
              <Td textAlign="center" key={`${i}-submit-at`}>
                {fmtDatetime(s.submittedAt!.toDate())}
              </Td>
              <Td textAlign="left">{s.taskTitle}</Td>
              <Td textAlign="center">{s.username}</Td>
              <Td textAlign="center">
                {langMetasBrief[s.langId as LangID]?.name ?? "不明"}
              </Td>
              <Td textAlign="center">{s.score}</Td>
              <Td textAlign="center">
                {judgeResult(s.id, s.judgeStatus)}
              </Td>
              <Td textAlign="center">
                {typeof s.execTimeMs === "number" ? `${s.execTimeMs} ms` : "-"}
              </Td>
              <Td textAlign="center">
                {typeof s.execMemoryKib === "number" ? `${s.execMemoryKib} KB` : "-"}
              </Td>
              <Td textAlign="center">
                <Link href={`/contests/${contestSlug}/submissions/${s.id}`}>詳細</Link>
              </Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
};

type ContestSubmissionListProps = {
  mode: "me" | "all";
  heading: string;
};

export const ContestSubmissionList = ({ mode, heading }: ContestSubmissionListProps) => {
  const contestSlug = useRouterContestSlug();
  const { contest } = useGetContest({ slug: contestSlug });
  const contestId = useMemo(() => contest?.id ?? 1, [contest]);

  const user = useAccessTokenClaimValue();

  const cardbody = (mode: "me" | "all") => {
    switch (mode) {
      case "me": {
        if (!user) {
          return (
            <Text>
              自分の提出を見るには<Link href="/login">ログイン</Link>してください
            </Text>
          );
        }
        return <SubmissionTable contestSlug={contestSlug} contestId={contestId} username={user.username} />;
      }
      case "all":
      default:
        return <SubmissionTable contestSlug={contestSlug} contestId={contestId} />;
    }
  };

  return (
    <Container h="100%" maxW="container.lg">
      <Card px={3} py={4} h="inherit" borderRadius={0}>
        <CardHeader>
          <Heading as="h1">{heading}</Heading>
        </CardHeader>
        <CardBody>
          {cardbody(mode)}
        </CardBody>
      </Card>
    </Container>
  );
};
