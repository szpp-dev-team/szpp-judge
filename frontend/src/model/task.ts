import { JudgeStatus, JudgeTestcaseProgress } from "@/src/model/judge";

export const DifficultyValues = [
  "beginner",
  "easy",
  "medium",
  "hard",
  "impossible",
] as const;

export type Difficulty = (typeof DifficultyValues)[number];

export type Task = {
  id: string;
  slug: string;
  difficulty: Difficulty;
  haiten: number;
  title: string;
};

type withSubmissionSummary = {
  status?: JudgeStatus;
  score?: number;
  progress?: JudgeTestcaseProgress;
};

export type TaskWithMySubmissionSummary = Task & withSubmissionSummary;
