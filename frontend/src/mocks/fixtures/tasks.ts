import { Difficulty, Task, Testcase } from "@/src/gen/proto/backend/v1/task_resources_pb";
import type { PlainMessage } from "@bufbuild/protobuf";
import { DUMMY_TASK1_STATEMENT } from "../fixtures/task1-statement";
import { DUMMY_TASK2_STATEMENT } from "../fixtures/task2-statement";

export const dummyTasks: Readonly<PlainMessage<Task>[]> = [
  {
    id: 1000,
    title: "マークダウン・KaTeXテスト1",
    execTimeLimit: 1000,
    execMemoryLimit: 256,
    difficulty: Difficulty.BEGINNER,
    statement: DUMMY_TASK1_STATEMENT,
  },
  {
    id: 1001,
    title: "マークダウン・KaTeXテスト2",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.EASY,
    statement: DUMMY_TASK2_STATEMENT,
  },
  {
    id: 1002,
    title: "すずっぴー君のおつかい",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.EASY,
    statement: "## 問題文\nすずっぴー君はおつかいに行きました。 ",
  },
  {
    id: 1003,
    title: "すずっぴー君の怒り",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.EASY,
    statement: "## 問題文\nすずっぴー君は激怒した。 ",
  },
  {
    id: 1004,
    title: "すずっぴー君の分裂",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.MEDIUM,
    statement: "## 問題文\nすずっぴー君は分裂を始めました。次はあなたの番です。 ",
  },
  {
    id: 1005,
    title: "すずっぴー君の爆散",
    execTimeLimit: 1000,
    execMemoryLimit: 512,
    difficulty: Difficulty.MEDIUM,
    statement: "## 問題文\nバーン。 ",
  },
  {
    id: 1006,
    title: "師匠との出会い",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.HARD,
    statement: "## 問題文\nにゃーん。 ",
  },
  {
    id: 1007,
    title: "師匠の爆散",
    execTimeLimit: 5000,
    execMemoryLimit: 2048,
    difficulty: Difficulty.HARD,
    statement: "## 問題文\nバーンバーンバーン。 ",
  },
  {
    id: 1008,
    title:
      "長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル",
    execTimeLimit: 500,
    execMemoryLimit: 256,
    difficulty: Difficulty.IMPOSSIBLE,
    statement: "## 問題文\nすずっぴー君はおつかいに行きました。 ",
  },
] as const;
