export type JudgeTestcaseProgress = {
  done: number;
  total: number;
};

export const JudgeStatusValues = [
  "WJ",
  "IE",
  "AC",
  "CE",
  "MLE",
  "OLE",
  "RE",
  "TLE",
  "WA",
] as const;

export type JudgeStatus = (typeof JudgeStatusValues)[number];

export const JudgeStatus = {
  toJapanese(s: JudgeStatus): string {
    return judgeStatusJaDict[s];
  },
  toEnglishAbbr(s: JudgeStatus): string {
    return s;
  },
} as const;

const judgeStatusJaDict = {
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
