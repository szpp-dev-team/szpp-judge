import Link from "next/link";

export const ContestList = () => {
  return (
    <div>
      <strong>動作確認用</strong>
      <ul>
        <li>
          <Link href="/contests/not_exist/tasks">
            Go to <code>note_exist/tasks</code>
          </Link>
        </li>
        <li>
          <Link href="/contests/sbc001/tasks">
            Go to <code>sbc001/tasks</code>
          </Link>
        </li>
      </ul>
    </div>
  );
};
