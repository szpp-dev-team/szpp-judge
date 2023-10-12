import { LangID } from "@/src/gen/langs";
import { cpp } from "@codemirror/lang-cpp";
import { java } from "@codemirror/lang-java";
import { python } from "@codemirror/lang-python";
import type { LanguageSupport } from "@codemirror/language";
import ReactCodeMirror, { ReactCodeMirrorProps } from "@uiw/react-codemirror";
import { useMemo } from "react";

type EditorProps = {
  lang?: LangID;
} & Omit<ReactCodeMirrorProps, "lang">;

const langExtension = {
  "c/11/gcc": cpp,
  "cpp/20/gcc": cpp,
  "scratch/3/gcc": cpp,
  "java/21/openjdk": java,
  "python/3.11/cpython": python,
} as const satisfies Record<LangID, () => LanguageSupport>;

export const Editor = ({ lang, extensions = [], ...props }: EditorProps) => {
  const exts = useMemo(
    () => lang == null ? extensions : [langExtension[lang](), ...extensions],
    [lang, extensions],
  );

  return <ReactCodeMirror extensions={exts} {...props} />;
};
