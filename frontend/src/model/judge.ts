export type JudgeTestcaseProgress = {
  done: number;
  total: number;
};

export type JudgeStatus =
  | "WJ"
  | "IE"
  | "AC"
  | "CE"
  | "MLE"
  | "OLE"
  | "RE"
  | "TLE"
  | "WA";

export const JudgeStatus = {
  toJapanese(s: JudgeStatus): string {
    return jugeStatusJaDict[s];
  },
  toEnglishAbbr(s: JudgeStatus): string {
    return s;
  },
} as const;

const jugeStatusJaDict = {
  WJ: "ジャッジ待ち",
  IE: "システム不具合",
  AC: "正解",
  CE: "コンパイルエラー",
  MLE: "メモリ制限エラー",
  OLE: "出力量制限エラー",
  RE: "実行時エラー",
  TLE: "時間制限エラー",
  WA: "出力値不正解",
} as const satisfies Record<JudgeStatus, string>;
