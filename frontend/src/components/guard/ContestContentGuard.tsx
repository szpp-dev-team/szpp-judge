import { useAccessTokenClaimValue } from "@/src/globalStates/credential";
import { useGetContest, useGetMyRegistrationStatus, useRouterContestSlug } from "@/src/usecases/contest";
import { getLoginUriWithRedirect } from "@/src/usecases/url";
import { Text } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { ReactNode } from "react";
import { CenteringFullSizeContainer } from "../ui/CenteringFullSizeContainer";
import { Link } from "../ui/Link";
import { LoadingPane } from "../ui/LoadingPane";

type ContestContentGuardProps = {
  /** 表示内容の名称。ユーザへの表示文で使う。 例: 問題一覧, 問題文, 提出結果 */
  contentName: string;
  children: ReactNode;
};

/**
 * コンテスト内容 (コンテスト時間中に参加者以外に見えては困る内容)  を適切に表示する。
 * URL から [contest_slug] を取得し、そのコンテストの終了・開始時刻、およびユーザが参加者かどうかで
 * children の表示の抑制や、ログイン・参加登録への誘導をする。
 */
export const ContestContentGuard = ({ contentName, children }: ContestContentGuardProps) => {
  const router = useRouter();
  const slug = useRouterContestSlug();
  const logined = useAccessTokenClaimValue() != null;
  const { contest } = useGetContest({ slug });
  const registrationStatus = useGetMyRegistrationStatus({ contestSlug: slug }, { enabled: logined });

  if (contest == null) {
    return <LoadingPane />;
  }

  const now = new Date();

  // コンテスト終了後はログイン/参加登録の是非に関係なく誰でも見れる
  if (contest.endAt!.toDate() < now) {
    return children;
  }

  // コンテスト中は参加登録者のみ見れる
  if (contest.startAt!.toDate() < now) {
    const message = `コンテスト時間中は参加登録をした人のみが${contentName}にアクセスできます。`;
    if (!logined) {
      return (
        <CenteringFullSizeContainer>
          <Text>
            {message}
            <br />まずは <Link href={getLoginUriWithRedirect(router.asPath)}>ログイン</Link> してください。
          </Text>
        </CenteringFullSizeContainer>
      );
    }
    if (registrationStatus.isLoading) {
      return <LoadingPane />;
    }
    if (registrationStatus.data?.registered) {
      return children;
    }
    return (
      <CenteringFullSizeContainer>
        <Text>
          {message}
          <br />参加登録は <Link href={`/contests/${slug}`}>コンテストトップページ</Link> でできます。
        </Text>
      </CenteringFullSizeContainer>
    );
  }

  return (
    <CenteringFullSizeContainer>
      <Text>コンテスト開始前なので内容を表示できません。</Text>
    </CenteringFullSizeContainer>
  );
};
