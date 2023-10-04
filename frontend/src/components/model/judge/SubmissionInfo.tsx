import { LangID, langMetasBrief } from "@/src/gen/langs";
import type { SubmissionDetail } from "@/src/gen/proto/backend/v1/judge_resources_pb";
import { JudgeStatus as PbJudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { JudgeStatus } from "@/src/model/judge";
import { Grid } from "@chakra-ui/react";
import type { GridProps } from "@chakra-ui/react";
import { JudgeStatusBadge } from "../../model/judge/JudgeStatusBadge";
import { Link } from "../../ui/Link";
import { Pair } from "../../ui/Pair";

export const SubmissionInfo = (props: GridProps & { submissionDetail?: SubmissionDetail }) => {
  const { submissionDetail, ...rest } = props;
  if (!submissionDetail) {
    return null;
  }

  return (
    <Grid
      {...rest}
      as="dl"
      textAlign="center"
      templateColumns="3fr 5fr"
      borderTop="1px"
      borderLeft="1px"
      borderColor="gray.300"
    >
      {/* sv-SE ローケルでのフォーマットは `YYYY-mm-dd HH:MM:SS`*/}
      {/* submittedAt は proto が optional じゃないので `!.` でいいはずだが念の為 `?.` */}
      <Pair data={["提出日時", submissionDetail.submittedAt?.toDate().toLocaleString("sv-SE")]} />
      <Pair data={["問題", <Link href="#" key={submissionDetail.id}>{submissionDetail.taskTitle}</Link>]} />
      <Pair
        data={[
          "ユーザ",
          <Link href={`/users/${submissionDetail.username}`} key={submissionDetail.username}>
            {submissionDetail.username}
          </Link>,
        ]}
      />
      <Pair data={["言語", langMetasBrief[submissionDetail.langId as LangID]?.name ?? "不明"]} />
      <Pair data={["得点", submissionDetail.score]} />
      <Pair
        data={[
          "結果",
          <JudgeStatusBadge
            key={submissionDetail.status}
            status={PbJudgeStatus[submissionDetail.status ?? 0].toString() as JudgeStatus}
          />,
        ]}
      />
      <Pair data={["実行時間", `${submissionDetail.execTimeMs} ms`]} />
      <Pair data={["メモリ", `${submissionDetail.execMemoryKib} KB`]} />
    </Grid>
  );
};
