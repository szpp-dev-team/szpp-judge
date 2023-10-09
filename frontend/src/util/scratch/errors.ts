type Sb3ConvertErrorKind =
  | "FailedExtractAsZip"
  | "ProjectJsonNotFound"
  | "MaybeSb2"
  | "EntryPointNotFound"
  | "MultipleEntryPoint"
  | "FailedConvertBlock";

const humanReadableMessages = {
  "FailedExtractAsZip": ["ZIPとしての展開に失敗しました", ""],
  "ProjectJsonNotFound": ["project.json が含まれていません", ""],
  "MaybeSb2": ["有効な Scratch 3.0 ファイルではありません", "もしかして Scratch <= 2.0 ですか？"],
  "EntryPointNotFound": [
    "プログラムの開始地点がありません",
    "「（はた）がクリックされたとき」ブロックを、配置してください！",
  ],
  "MultipleEntryPoint": [
    "エラー: プログラムの開始地点が 2 つ以上あります",
    "「（はた）がクリックされたとき」ブロックは、1 つだけにしてください！",
  ],
  "FailedConvertBlock": ["以下のブロックを変換中にエラーが発生しました", ""],
} as const satisfies Record<Sb3ConvertErrorKind, readonly [string, string]>;

export class Sb3ConvertError extends Error {
  kind: Sb3ConvertErrorKind;
  detail: string;

  static {
    this.prototype.name = "Sb3ConvertError";
  }

  constructor(fileName: string, kind: Sb3ConvertErrorKind, detail: string = "", opt?: ErrorOptions) {
    super(humanReadableMessages[kind][0], opt);
    this.kind = kind;
    this.detail = "ファイル名: " + fileName;

    const s = humanReadableMessages[kind][1] + detail;
    if (s) {
      this.detail += ": ";
      this.detail += s;
    }
  }
}
