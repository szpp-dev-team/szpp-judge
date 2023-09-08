export const DifficultyValues = [
  "beginner",
  "easy",
  "medium",
  "hard",
  "impossible",
] as const;

export type Difficulty = (typeof DifficultyValues)[number];
