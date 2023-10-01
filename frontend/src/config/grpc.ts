import { AuthService } from "@/src/gen/proto/backend/v1/auth_service-AuthService_connectquery";
import type { Credential } from "@/src/model/auth";
import { Code, ConnectError, Interceptor } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

export const backendBaseUrl = process.env.NEXT_PUBLIC_BACKEND_BASE_URL!;

export const backendTransportWithOnlyBaseUrl = createConnectTransport({
  baseUrl: backendBaseUrl,
});

const LOGIN_URL_PATH = AuthService.typeName + "/" + AuthService.methods.login.name;
const REFRESH_ACCESS_TOKEN_URL_PATH = AuthService.typeName + "/" + AuthService.methods.refreshAccessToken.name;

type AuthInterceptorOptions = {
  cred: Readonly<Credential>;
  accessTokenExpireAt?: number;
  refreshAccessToken: () => Promise<string>;
};

/**
 * - リクエストヘッダに `Authorization: Bearer {accessToken}` をセットする。
 * - リクエストした結果、accessToken の期限切れだぞというレスポンスが返ってきた場合は
 *   `refreshAccessToken` を呼び出す。
 */
const authInterceptor = ({
  cred: { accessToken, refreshToken },
  accessTokenExpireAt,
  refreshAccessToken,
}: AuthInterceptorOptions): Interceptor =>
(next) =>
async (req) => {
  if (!accessToken || req.url.endsWith(LOGIN_URL_PATH) || req.url.endsWith(REFRESH_ACCESS_TOKEN_URL_PATH)) {
    return await next(req);
  }

  if (accessTokenExpireAt && Date.now() > accessTokenExpireAt) {
    const newAccessToken = await refreshAccessToken();
    req.header.set("Authorization", `Bearer ${newAccessToken}`);
  } else {
    req.header.set("Authorization", `Bearer ${accessToken}`);
  }

  try {
    return await next(req);
  } catch (e) {
    if (e instanceof ConnectError) {
      // アクセストークンの期限切れなのでリフレッシュしてから再試行
      // ログイン失敗時にも Unauthenticated が返ってくるが、上の if 文で early return しているので大丈夫なはず
      if (e.code === Code.Unauthenticated && refreshToken) {
        const newAccessToken = await refreshAccessToken();
        req.header.set("Authorization", `Bearer ${newAccessToken}`);
        return await next(req);
      }
    }
    throw e;
  }
};

export const createBackendTransport = (
  authInterceptorOpt: AuthInterceptorOptions,
) =>
  createConnectTransport({
    baseUrl: backendBaseUrl,
    interceptors: [authInterceptor(authInterceptorOpt)],
  });
