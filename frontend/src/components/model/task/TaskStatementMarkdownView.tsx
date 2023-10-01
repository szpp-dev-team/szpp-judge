import { parseMarkdownToHtml } from "@/src/util/markdown";
import { Box, type BoxProps } from "@chakra-ui/react";
import { useMemo } from "react";
import style from "./TaskStatementMarkdownView.module.scss";
import "katex/dist/katex.min.css";

export type TaskStatementMarkdownViewProps = {
  markdown?: string;
} & Omit<BoxProps, "children">;

export const TaskStatementMarkdownView = ({
  markdown,
  ...props
}: TaskStatementMarkdownViewProps) => {
  console.log(markdown);
  const html = useMemo(() => parseMarkdownToHtml(markdown ?? ""), [markdown]);

  return (
    <Box
      className={style.TaskStatementMarkdownView}
      {...props}
      dangerouslySetInnerHTML={{ __html: html }}
    />
  );
};
