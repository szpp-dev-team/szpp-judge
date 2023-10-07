import { parseMarkdownToHtml } from "@/src/util/markdown";
import { useMemo } from "react";
import "katex/dist/katex.min.css";
import { Box, type BoxProps } from "@chakra-ui/react";

export type MarkdownViewProps = {
  markdown?: string;
} & Omit<BoxProps, "children">;

export const MarkdownView = ({
  markdown,
  ...props
}: MarkdownViewProps) => {
  const html = useMemo(() => parseMarkdownToHtml(markdown ?? ""), [markdown]);

  return <Box {...props} dangerouslySetInnerHTML={{ __html: html }} />;
};
