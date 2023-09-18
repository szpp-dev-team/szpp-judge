const CHAR_CODE_A = "A".charCodeAt(0);

/**
 * コンテストの問題数に応じて問題番号をアルファベットまたは整数で生成する。
 * @param i 問題の添字 (0-indexed)
 * @param numTotalTasks コンテストの問題数
 */
export const genNthTaskSeq = (i: number, numTotalTasks: number): string => {
  if (numTotalTasks < 26) {
    return String.fromCharCode(CHAR_CODE_A + i);
  }
  return String(i + 1);
};
