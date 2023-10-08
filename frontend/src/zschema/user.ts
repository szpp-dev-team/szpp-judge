import { defaultErrorMap } from "@/src/config/zod";
import { z } from "zod";

// _app.tsx で z.setErrorMap する場合 zod が不要なページでも zod を import することになり、
// バンドルサイズが大きくなってしまうので個々の zschema/*.ts で読み込むようにしている。
z.setErrorMap(defaultErrorMap);

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
  .regex(/^[^ ]*$/, "半角スペースが含まれています")
  .regex(/^[!-~]*$/, "全角文字が含まれています"); // パスワードを表示するボタンを押すと全角文字が入力できてしまうのでチェック

const registrationPassword = loginPassword
  .describe("半角英字・数字ともに必須、8文字以上60文字以内")
  .min(8)
  .regex(/[A-Za-z]/, "半角英字を含めてください")
  .regex(/[0-9]/, "半角数字を含めてください");

const email = z.string()
  .min(1)
  .email();

export const userLoginSchema = z.object({
  username,
  password: loginPassword,
});

export const userRegistrationSchema = z.object({
  email,
  username,
  password: registrationPassword,
  confPassword: loginPassword,
}).superRefine((values, ctx) => {
  if (values.password !== values.confPassword) {
    ctx.addIssue({
      path: ["confPassword"],
      code: "custom",
      message: "パスワードが一致しません",
    });
  }
});
