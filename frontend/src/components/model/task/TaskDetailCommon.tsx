import style from "@/src/components/model/task/TaskDetailCommon.module.scss";
import { Task } from "@/src/gen/proto/backend/v1/task_resources_pb";
import { Difficulty } from "@/src/model/task";
import { Box, Card, Heading, Text } from "@chakra-ui/react";
import dynamic from "next/dynamic";
import { DifficultyBadge } from "../../model/task/DifficultyBadge";

export type TaskDetailCommonProps = {
  task: Task;
  seqCode: string;
  score?: number;
};

// コンテスト問題ページ、スタンドアロン問題ページ共通の問題表示ビュー。
// 問題タイトルや時間制限、問題文、サンプルケース、提出フォームなどを表示する
export const TaskDetailCommon = ({
  task,
  seqCode,
  score,
}: TaskDetailCommonProps) => {
  const MarkdownView = dynamic(
    () => import("@/src/components/ui/MarkdownView").then(mod => mod.MarkdownView),
    { ssr: false },
  );
  return (
    <Card px={6} py={4} h="100%" rounded={"none"}>
      <Heading as="h1">{seqCode} - {task.title}</Heading>
      <Box>
        {/* TODO: 良いカンジに表示する */}
        <Text>実行時間制限:</Text>
        <Text>{task.execTimeLimit / 1000} ms</Text>
        <Text>メモリ制限:</Text>
        <Text>{task.execMemoryLimit} MiB</Text>
        <Text>難易度:</Text>
        <DifficultyBadge dif={Difficulty.fromPb(task.difficulty)} />
      </Box>
      <Box className={style.markdownWrapper}>
        <MarkdownView markdown={task.statement} />
        <h2>入出力例</h2>
      </Box>
    </Card>
  );
};
