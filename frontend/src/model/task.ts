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

export const ScoreStatusValues = [
  undefined,
  "perfect",
  "partial",
  "zero",
] as const;

/**
 * - undefined: 未提出
 * - "perfect": 満点
 * - "partial": 部分点
 * - "zero": 0
 */
export type ScoreStatus = (typeof ScoreStatusValues)[number];
