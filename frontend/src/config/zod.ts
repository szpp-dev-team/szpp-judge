import { ZodErrorMap, ZodIssueCode } from "zod";

export const defaultErrorMap: ZodErrorMap = (issue, ctx) => {
  const defaultMessage = { message: ctx.defaultError };

  switch (issue.code) {
    case ZodIssueCode.invalid_type:
      if (issue.received === "undefined") {
        return { message: "入力が空です" };
      }
      return defaultMessage;

    case ZodIssueCode.too_big: {
      if (issue.type === "string") {
        return {
          message: `${issue.maximum}文字${issue.inclusive ? "以内" : "未満"}にしてください`,
        };
      }
      if (issue.type === "array") {
        return {
          message: `選択数は${issue.maximum}つ${issue.inclusive ? "以内" : "未満"}にしてください`,
        };
      }
      return defaultMessage;
    }

    case ZodIssueCode.too_small: {
      if (issue.type === "string") {
        return {
          message: issue.minimum === 1
            ? "入力が空です"
            : `${issue.minimum}文字${issue.inclusive ? "以上に" : "より長く"}してください`,
        };
      }
      if (issue.type === "array") {
        return {
          message: `${issue.minimum}つ${issue.inclusive ? "以上" : "より多く"}選択してください`,
        };
      }
      return defaultMessage;
    }
  }
  return defaultMessage;
};
