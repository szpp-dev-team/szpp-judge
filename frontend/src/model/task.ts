export const DifficultyValues = [
  "beginner",
  "easy",
  "medium",
  "hard",
  "impossible",
] as const;

export type Difficulty = (typeof DifficultyValues)[number];

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
