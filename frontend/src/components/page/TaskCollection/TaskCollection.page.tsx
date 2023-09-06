import { TaskWithMySubmissionSummary } from "@/src/model/task";
import { pageTitle } from "@/src/util/pagemeta";
import Head from "next/head";
import { WithHeaderFooter } from "../../layout/WithHeaderFooter";
import { TaskCollection } from "./TaskCollection";

const tasks: TaskWithMySubmissionSummary[] = [
  {
    id: "A",
    difficulty: "beginner",
    haiten: 100,
    title: "すずっぴー君のおつかい",
    status: "AC",
    score: 100,
  },
  {
    id: "B",
    difficulty: "beginner",
    haiten: 200,
    title: "すずっぴー君の怒り",
    status: "WA",
    score: 0,
  },
  {
    id: "C",
    difficulty: "easy",
    haiten: 300,
    title: "すずっぴー君の分裂",
    status: "AC",
    score: 300,
  },
  {
    id: "D",
    difficulty: "easy",
    haiten: 314,
    title: "すずっぴー君の爆散",
  },
  {
    id: "E",
    difficulty: "medium",
    haiten: 400,
    title: "師匠との出会い",
    status: "WJ",
    progress: {
      done: 12,
      total: 40,
    },
  },
  {
    id: "F",
    difficulty: "beginner",
    haiten: 500,
    title: "師匠の爆散",
    status: "TLE",
    score: 50,
  },
  {
    id: "G",
    difficulty: "hard",
    haiten: 600,
    title: "すずっぴー君の凝固",
    status: "RE",
    score: 0,
  },
  {
    id: "Ex",
    difficulty: "impossible",
    haiten: 1333,
    title:
      "長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル",
  },
];

export const TaskCollectionPage = () => {
  return (
    <WithHeaderFooter>
      <Head>
        <title>{pageTitle("問題")}</title>
      </Head>
      <TaskCollection tasks={tasks} />
    </WithHeaderFooter>
  );
};
