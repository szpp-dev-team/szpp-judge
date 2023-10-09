import { indentWithTab } from "@codemirror/commands";
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
  keymap,
  lineNumbers,
} from "@codemirror/view";
import { useEffect, useRef } from "react";

interface EditorProps extends Omit<EditorStateConfig, "extensions"> {
  /** 編集できなくなる */
  readonly?: boolean;

  onDocChange?: (s: string) => unknown;

  height?: string;
}

export const Editor = ({
  doc,
  height,
  onDocChange = () => {},
  readonly = false,
}: EditorProps) => {
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
        keymap.of([indentWithTab]),
        cpp(),
        EditorState.readOnly.of(readonly),
        EditorView.updateListener.of(upd => {
          if (upd.docChanged) {
            onDocChange(upd.state.doc.toString());
          }
        }),
      ],
    });

    const view = new EditorView({
      state,
      parent: editor.current,
    });

    return () => {
      view.destroy();
    };
  }, [doc, readonly]);

  return <div style={{ height, overflow: "auto", overscrollBehavior: "none" }} ref={editor} />;
};
