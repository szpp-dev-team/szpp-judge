import { Code, ConnectError, createPromiseClient } from "@connectrpc/connect";
import { useMutation } from "@tanstack/react-query";
import { backendTransportWithOnlyBaseUrl } from "../config/connectRpc";
import { AuthService, login, logout } from "../gen/proto/backend/v1/auth_service-AuthService_connectquery";
import { RefreshAccessTokenResponse } from "../gen/proto/backend/v1/auth_service_pb";
import {
  useCredentialSetter,
  useCredentialValueAndEraser,
  useCredentialValueAndSetter,
} from "../globalStates/credential";
import { Duration } from "../util/time";

/**
 * ログイン用カスタムフック
 * - フックを呼び出す側は mutate を呼び出す
 * - mutate が成功したらこのフックによって状態( jotai )が更新されるのでフックを呼び出した側は何もすべきでない
 * @example
 * ```ts
 * const { mutate } = useLogin(() => { toast({ title: "失敗" }}) });
 * const onSubmit = handleSubmit((values) => mutate(values));
 * ```
 * @param onUnauthenticatedError パスワード間違いをコンポーネントに通知するためのコールバック
 * @returns
 */
export const useLogin = (onUnauthenticatedError?: () => void) => {
  const onConnectError = (e: ConnectError) => {
    if (e.code === Code.Unauthenticated && onUnauthenticatedError) {
      onUnauthenticatedError();
    } else {
      throw e; // mutate を使う限り React Query がエラーをキャッチして(握りつぶして)くれる
    }
  };

  const setCredential = useCredentialSetter();
  const { data, error, isLoading, mutate } = useMutation({
    ...login.useMutation({ onError: onConnectError }),
    onSuccess: (data) => {
      console.log("[useLoginHook] onSuccess");

      setCredential({
        accessToken: data.accessToken,
        refreshToken: data.refreshToken,
      });
    },
  });

  return { data, error, isLoading, mutate };
};

export const useLogout = () => {
  const [cred, eraseCred] = useCredentialValueAndEraser();

  const { isLoading, mutate: mutateImpl } = useMutation({
    ...logout.useMutation(),
    onSettled: () => {
      eraseCred();
    },
  });

  const mutate = (options?: Parameters<typeof mutateImpl>[1]) => {
    return mutateImpl({ refreshToken: cred.refreshToken }, options);
  };

  return { isLoading, mutate };
};

/**
 * Tanstack の useQuery を使わずに rpc AuthService.refreshAccessToken を呼び出す。
 * リクエスト送信後、 STALE_TIME 時間の間は結果をキャッシュする。
 *
 * # なぜ useQuery を使わない実装か？
 * authInterceptor 作成時 (in _app.tsx) は QueryClientProvider の外なので useQuery を呼び出すとエラーになるため。
 */
export const useRefreshAccessTokenWithoutQueryClient = () => {
  const STALE_TIME = Duration.SECOND * 3;
  const [cred, setCred] = useCredentialValueAndSetter();

  // リアクティブにする必要はない & useState の setXXX が反映されるのは次のレンダリング時であって即時ではない
  let lastReqAt = new Date(0);
  let response: null | Promise<RefreshAccessTokenResponse> = null;

  const cli = createPromiseClient(AuthService, backendTransportWithOnlyBaseUrl);

  const refreshAccessToken = async (): Promise<RefreshAccessTokenResponse> => {
    const now = new Date();
    if (now.getTime() < lastReqAt.getTime() + STALE_TIME && response != null) {
      return await response;
    }
    lastReqAt = now;

    response = cli.refreshAccessToken({ refreshToken: cred.refreshToken });

    const resp = await response;
    setCred({ ...cred, accessToken: resp.accessToken });

    return resp;
  };

  return { refreshAccessToken };
};
