import JSZip from "jszip";
import { Sb3ToCppConverter } from "./cpp_converter";
import { Sb3ConvertError } from "./errors";

const extractProjectJsonFromSb3 = async (file: File): Promise<string> => {
  let zipFile: JSZip;
  try {
    zipFile = await JSZip.loadAsync(file);
  } catch (e) {
    throw new Sb3ConvertError(file.name, "FailedExtractAsZip", "", { cause: e });
  }

  const jsonFile = zipFile.files["project.json"];
  if (jsonFile == null) {
    throw new Sb3ConvertError(file.name, "ProjectJsonNotFound");
  }

  return jsonFile.async("text");
};

export const convertSb3ToCpp = async (file: File) => {
  const projectJsonContent = await extractProjectJsonFromSb3(file);
  return convertSb3JsonToCpp(file.name, projectJsonContent);
};

/**
 *  参考: https://github.com/yos1up/scratch2cpp/blob/eaa59832f752fbbf665b77e59567a9eec6670c21/web/crx/sb3_to_cpp.js
 *
 *  変換に問題があった場合、
 *  CPPソースへの変換ができないレベルのエラーでは Sb3ConvertError を投げ、
 *  そうでなければ 戻り値の warnings に警告をセットする。
 *
 *  @throws Sb3ConvertError  致命的なエラーの場合
 */
const convertSb3JsonToCpp = (
  fileName: string,
  projectJsonContent: string,
): { cppSource: string; warnings: string[] } => {
  const converter = new Sb3ToCppConverter();
  const rslt = converter.convert(projectJsonContent);
  const cppSource = rslt[0]! as string;
  const errorInfos = rslt[1]!;

  // errorInfos が string になることはないと思われる
  if (typeof errorInfos === "string") {
    throw new Error();
  }

  const warnings: string[] = [];

  for (const errorInfo of errorInfos) {
    switch (errorInfo["code"]) {
      case -1:
        throw new Sb3ConvertError(fileName, "MaybeSb2");
      case 1:
        warnings.push("次のプロックは変換できなかったので無視しました: " + errorInfo.message);
        break;
      case 2:
        throw new Sb3ConvertError(fileName, "EntryPointNotFound");
      case 3:
        throw new Sb3ConvertError(fileName, "MultipleEntryPoint");
      case 4:
        // エラー番号 4 は使われていない (cpp_converter.js でコメントアウトされているので無視)

        // errorMessage = {
        //   'en': 'CAUTION: Non-ascii characters are used in identifiers. Please select Clang compiler (C++14 (Clang 3.8.0)) to compile it correctly.',
        //   'ja': 'ちゅうい: へんすうめいに アスキーもじ いがいが ふくまれるため、Clang コンパイラ (C++14 (Clang 3.8.0)) をせんたくしてください！'
        // }[lang];
        break;

      case 5:
        throw new Sb3ConvertError(fileName, "FailedConvertBlock", errorInfo.message);
    }
  }

  return { cppSource, warnings };
};
