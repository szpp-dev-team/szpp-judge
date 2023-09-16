import { Code, ConnectError } from "@bufbuild/connect";
import { useMutation } from "@tanstack/react-query";
import { login } from "../gen/proto/backend/v1/services-AuthService_connectquery";

/**
 * - 使うときは mutate を呼び出す
 * - mutate を使う限りエラーは react-query で握りつぶしてくれるが username, password の
 *   間違いが原因のエラーに関してはコンポーネントに通知したいのでコールバックがある
 * @param onUnauthenticatedError
 * @returns
 */
export const useLogin = (onUnauthenticatedError?: () => void) => {
  const onConnectError = (e: ConnectError) => {
    if (e.code === Code.Unauthenticated && onUnauthenticatedError) {
      onUnauthenticatedError();
    } else {
      throw e;
    }
  };

  const { data, error, isLoading, mutate, mutateAsync } = useMutation({
    ...login.useMutation({ onError: onConnectError }),
    onSuccess: (data) => {
      console.log("[usecase]onSuccess");
    },
  });

  return { data, error, isLoading, mutate, mutateAsync };
};
