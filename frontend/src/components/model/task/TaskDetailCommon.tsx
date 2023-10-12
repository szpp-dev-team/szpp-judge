import taskDetailStyle from "@/src/components/model/task/TaskDetailCommon.module.scss";
import { MAX_SOURCE_CODE_SIZE } from "@/src/config/submission";
import { LangID } from "@/src/gen/langs";
import { SubmitRequest } from "@/src/gen/proto/backend/v1/judge_service_pb";
import { Task, Testcase } from "@/src/gen/proto/backend/v1/task_resources_pb";
import { useSubmissionLangSetter } from "@/src/globalStates/submission";
import { Difficulty } from "@/src/model/task";
import { useSubmit } from "@/src/usecases/judge";
import { PlainMessage } from "@bufbuild/protobuf";
import { Box, BoxProps, Card, Heading, Text, useToast } from "@chakra-ui/react";
import dynamic from "next/dynamic";
import { ReactNode, useCallback } from "react";
import { DifficultyBadge } from "../../model/task/DifficultyBadge";
import { TestcaseView } from "./TestcaseView";

export type TaskDetailCommonProps = {
  task: Task;
  sampleCases: Testcase[];
  onSubmitSuccess: (submissionId: number) => void;
  contestId?: number;
  taskSeqCode?: string;
  score?: number;
};

const MarkdownView = dynamic(
  () => import("@/src/components/ui/MarkdownView").then(mod => mod.MarkdownView),
  {
    loading: () => <p>読み込み中です</p>,
    ssr: false,
  },
);

const SubmissionEditor = dynamic(
  () => import("@/src/components/model/judge/SubmissionForm").then(mod => mod.SubmissionForm),
  {
    loading: () => <p>読み込み中です</p>,
    ssr: false,
  },
);

// コンテスト問題ページ、スタンドアロン問題ページ共通の問題表示ビュー。
// 問題タイトルや時間制限、問題文、サンプルケース、提出フォームなどを表示する
export const TaskDetailCommon = ({
  task,
  sampleCases,
  onSubmitSuccess,
  contestId,
  taskSeqCode,
  score,
}: TaskDetailCommonProps) => {
  const toast = useToast();

  const { mutate, isLoading: isSubmissionLoading } = useSubmit();
  const saveSubmissionLang = useSubmissionLangSetter();

  const handleSubmit = useCallback((langId: LangID, sourceCode: string) => {
    if (sourceCode.length === 0) {
      toast({
        title: "ソースコードが空です。",
        status: "error",
        duration: 1000,
      });
      return;
    }

    saveSubmissionLang(langId);

    if (sourceCode.length > MAX_SOURCE_CODE_SIZE) {
      toast({
        title: `ソースコード長の上限は ${MAX_SOURCE_CODE_SIZE} KiBです。`,
        description: `あなたのソースコードの長さ: ${sourceCode.length} Byte (${sourceCode.length >> 10} KiB)`,
        status: "error",
        duration: 5000,
      });
      return;
    }

    const req: PlainMessage<SubmitRequest> = {
      contestId,
      taskId: task.id,
      langId,
      sourceCode,
    };
    mutate(req, {
      onSuccess: ({ submissionId }) => {
        onSubmitSuccess(submissionId);
      },
      onError: (e) => {
        console.error(e);
        toast({
          title: "提出に失敗しました",
          description: e.message,
          status: "error",
          duration: 5000,
        });
      },
    });
  }, [contestId, task.id, mutate, onSubmitSuccess, toast, saveSubmissionLang]);

  return (
    <Card px={6} py={4} minH="100%" maxW="860px" w="100%" rounded={"none"} color="cyan.900">
      <Heading as="h1" fontSize="3xl">
        {taskSeqCode && taskSeqCode + " - "}
        {task.title}
      </Heading>
      <Box mt={4} as="dl">
        <TaskMetaRecord name="実行時間制限">{task.execTimeLimit / 1000} sec</TaskMetaRecord>
        <TaskMetaRecord name="メモリ制限">{task.execMemoryLimit} MiB</TaskMetaRecord>
        <TaskMetaRecord name="難易度">
          <DifficultyBadge dif={Difficulty.fromPb(task.difficulty)} />
        </TaskMetaRecord>
      </Box>
      {score != null && (
        <Box mt={1} as="dl">
          <TaskMetaRecord name="配点" fontSize="lg">
            <Text as="strong">{score}</Text> 点
          </TaskMetaRecord>
        </Box>
      )}
      <MarkdownView markdown={task.statement} className={taskDetailStyle.markdownWrapper} mt={10} />
      <h2 className={taskDetailStyle.h2}>入出力例</h2>
      {sampleCases.map((c, i) => (
        <Box key={c.id} as="section" mb={8}>
          <h3 className={taskDetailStyle.h3}>入出力例 {i + 1}</h3>
          <TestcaseView title={`入力例 ${i + 1}`} titleAs="h4" content={c.input} my={2} />
          <TestcaseView title={`出力例 ${i + 1}`} titleAs="h4" content={c.output} my={4} />
          <MarkdownView markdown={c.description} className={taskDetailStyle.markdownWrapper} />
        </Box>
      ))}
      <h2 className={taskDetailStyle.h2}>解答プログラム</h2>
      <SubmissionEditor
        onSubmit={handleSubmit}
        submitButtonProps={{ isLoading: isSubmissionLoading }}
      />
    </Card>
  );
};

const TaskMetaRecord = ({ name, children, ...props }: {
  name: string;
  children: ReactNode;
} & Omit<BoxProps, "name" | "children">) => {
  return (
    <Box
      display="inline-flex"
      alignItems="center"
      sx={{
        "&:not(:last-child)::after": {
          content: `'/'`,
          mx: 3,
        },
      }}
      {...props}
    >
      <Text
        as="dt"
        display="inline"
        _after={{ content: `':'`, mr: 1 }}
      >
        {name}
      </Text>
      <Box as="dd" display="inline">{children}</Box>
    </Box>
  );
};
