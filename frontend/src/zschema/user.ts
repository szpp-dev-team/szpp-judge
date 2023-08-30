import { z } from "zod";

const username = z.string()
  .describe("半角英数字・アンダーバー・ハイフンのみ使用可、20文字以内")
  .min(1)
  .max(20)
  .regex(
    /^[a-zA-Z0-9_-]+$/,
    "使用可能な文字は半角英数字・アンダーバー・ハイフンのみです",
  );

// Go 言語の bcrypt パッケージ (パスワードハッシュライブラリ)
// が対応する最大文字数が72文字なので、ソルトの追加も考慮して最大文字数は60文字とする
const loginPassword = z.string()
  .min(1)
  .max(60)
  .regex(/^[^\x00-\x1F\x80-\xFF]+$/, "不正な文字が含まれています");

const registrationPassword = loginPassword
  .describe("半角英大文字・英小文字・数字・記号全て必須、8文字以上60文字以内")
  .min(8)
  .regex(/[A-Z]/, "半角英大文字を含めてください")
  .regex(/[a-z]/, "半角英小文字を含めてください")
  .regex(/[0-9]/, "半角数字を含めてください")
  .regex(/[\x20-\x2F\x3A-\x40\x5B-\x60\x7B-\x7E]/, "半角記号を含めてください");

const email = z.string()
  .email();

export const userLoginSchema = z.object({
  username,
  password: loginPassword,
});

export const userRegistrationSchema = z.object({
  email,
  username,
  password: registrationPassword,
});
