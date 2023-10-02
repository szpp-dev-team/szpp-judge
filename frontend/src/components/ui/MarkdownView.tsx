import { parseMarkdownToHtml } from "@/src/util/markdown";
import { useMemo } from "react";
import "katex/dist/katex.min.css";

export type MarkdownViewProps = {
  markdown?: string;
};

export const MarkdownView = ({
  markdown,
}: MarkdownViewProps) => {
  const html = useMemo(() => parseMarkdownToHtml(markdown ?? ""), [markdown]);

  return <div dangerouslySetInnerHTML={{ __html: html }} />;
};
