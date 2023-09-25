import { Link } from "@/src/components/ui/Link";
import { useListContests } from "@/src/usecases/contest";

export const ContestList = () => {
  const { contests } = useListContests({});
  return (
    <div>
      <strong>動作確認用</strong>
      <ul>
        <li>
          <Link href="/contests/not_exist/tasks">
            Go to <code>note_exist/tasks</code>
          </Link>
        </li>
        {contests?.map((c) => (
          <li key={c.id}>
            <Link href={`/contests/${c.slug}`}>Go to {c.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};
