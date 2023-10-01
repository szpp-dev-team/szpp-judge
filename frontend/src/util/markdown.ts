import DOMPurify from "dompurify";
import katex from "katex";
import { marked } from "marked";

const ESCAPE_SET = /\\[#$%&,.<>:;{}_=^`'"\\]/g;

export const parseMarkdownToHtml = (s: string): string => {
  s = s.replaceAll(ESCAPE_SET, (capture0) => "\\" + capture0);
  s = marked.use({ gfm: true }).parse(s);
  s = replaceKatexToMathHtml(s);
  s = s.replaceAll("\\$", "$");
  s = DOMPurify.sanitize(s);
  return s;
};

const KATEX_RANGE = /(?<=^|[^\\])(\${1,2})(?!\$)((?:[^])*?[^\\$])\1(?!\$)/g;

const replaceKatexToMathHtml = (s: string): string => {
  return s.replaceAll(KATEX_RANGE, (_, capture1, capture2) => {
    const formula = capture2
      .replaceAll("&amp;", "&")
      .replaceAll("&quot;", "\"")
      .replaceAll("&gt;", ">")
      .replaceAll("&lt;", "<")
      .replaceAll("\\<br>", "\\\\\n");
    return katex.renderToString(formula, {
      displayMode: capture1.length === 2,
      throwOnError: false,
    });
  });
};
