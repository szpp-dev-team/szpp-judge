import { cpp } from "@codemirror/lang-cpp";
import { defaultHighlightStyle, syntaxHighlighting } from "@codemirror/language";
import { EditorState, EditorStateConfig } from "@codemirror/state";
import {
  drawSelection,
  dropCursor,
  EditorView,
  highlightActiveLine,
  highlightActiveLineGutter,
  highlightSpecialChars,
  lineNumbers,
} from "@codemirror/view";
import { useEffect, useRef } from "react";

interface EditorProps extends Omit<EditorStateConfig, "extensions"> {
  /** 編集できなくなる */
  readonly?: boolean;
}

export const Editor = ({ doc, readonly = false }: EditorProps) => {
  const editor = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!editor.current) return;

    const state = EditorState.create({
      doc,
      extensions: [
        lineNumbers(),
        highlightActiveLineGutter(),
        highlightSpecialChars(),
        drawSelection(),
        dropCursor(),
        syntaxHighlighting(defaultHighlightStyle, { fallback: true }),
        highlightActiveLine(),
        cpp(),
        EditorState.readOnly.of(readonly),
      ],
    });

    const view = new EditorView({
      state,
      parent: editor.current,
    });

    return () => {
      view.destroy();
    };
  }, [editor.current]);

  return <div ref={editor} />;
};
