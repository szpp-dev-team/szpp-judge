export const DifficultyValues = [
  "beginner",
  "easy",
  "medium",
  "hard",
  "impossible",
] as const;

export type Difficulty = (typeof DifficultyValues)[number];

const judgeStatuses = [
  "ジャッジ待ち",
  "正解",
  "コンパイルエラー",
  "メモリ制限エラー",
  "出力量制限エラー",
  "実行時エラー",
  "時間制限エラー",
  "出力値不正解",
  "システム不具合",
] as const;
type JudgeStatus = (typeof judgeStatuses)[number];

export type Task = {
  id: string;
  difficulty: Difficulty;
  haiten: number;
  title: string;
};

type withSubmissionSummary = {
  status?: JudgeStatus;
  score?: number;
};

export type TaskWithMySubmissionSummary = Task & withSubmissionSummary;
