import { Difficulty as PbDifficulty } from "@/src/gen/proto/backend/v1/task_resources_pb";

export const DifficultyValues = [
  "beginner",
  "easy",
  "medium",
  "hard",
  "impossible",
] as const;

export type Difficulty = (typeof DifficultyValues)[number];

export const Difficulty = {
  fromPb(d: PbDifficulty): Difficulty {
    switch (d) {
      case PbDifficulty.DIFFICULTY_UNSPECIFIED: {
        throw new Error(`Got a DIFFICULTY_UNSPECIFIED`);
      }
      case PbDifficulty.BEGINNER:
        return "beginner";
      case PbDifficulty.EASY:
        return "easy";
      case PbDifficulty.MEDIUM:
        return "medium";
      case PbDifficulty.HARD:
        return "hard";
      case PbDifficulty.IMPOSSIBLE:
        return "impossible";
      default: {
        const exhaustiveCheck: never = d;
        throw new Error(`Invlid pb difficulty: ${exhaustiveCheck}`);
      }
    }
  },
};

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

export const ScoreStatus = {
  fromScore: (allotScore: number, userScore: number | undefined): ScoreStatus => {
    if (userScore == null) return undefined;
    if (userScore === 0) return "zero";
    if (userScore < allotScore) return "partial";
    return "perfect";
  },
};
