import { Difficulty as PbDifficulty } from "@/src/gen/proto/backend/v1/task_resources_pb";
import { Difficulty } from "../task";
import { Converter } from "./type";

const difficultyConverter: Converter<Difficulty, PbDifficulty> = {
  from: (dif) => {
    switch (dif) {
      case PbDifficulty.DIFFICULTY_UNSPECIFIED:
        throw new Error("Unspecified difficulty cannot be handled");
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
        const exhaustiveCheck: never = dif;
        throw new Error(`Invlid difficulty: ${exhaustiveCheck}`);
      }
    }
  },
  to: (dif) => {
    switch (dif) {
      case "beginner":
        return PbDifficulty.BEGINNER;
      case "easy":
        return PbDifficulty.EASY;
      case "medium":
        return PbDifficulty.MEDIUM;
      case "hard":
        return PbDifficulty.HARD;
      case "impossible":
        return PbDifficulty.IMPOSSIBLE;
      default:
        return PbDifficulty.DIFFICULTY_UNSPECIFIED;
    }
  },
};

export const taskConverter = {
  difficulty: difficultyConverter,
};
