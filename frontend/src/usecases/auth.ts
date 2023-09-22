import { Code, ConnectError } from "@bufbuild/connect";
import { useMutation } from "@tanstack/react-query";
import { login, logout } from "../gen/proto/backend/v1/auth_service-AuthService_connectquery";
import { useCredentialSetter, useCredentialValueAndEraser } from "../globalStates/credential";

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
