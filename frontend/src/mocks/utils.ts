import { decodeJwtPayload } from "@/src/util/jwt";

/**
 * Authorization ヘッダの値を JWT としてパースし、
 * トークンの発行時刻と `tokenLifeMilliSec` から期限切れか否かを判定する。
 *
 * 使用例:
 * ```ts
 * grpcMock(..., async (ctx, res, decodeReq, encodeResp, req) => {
 *    // 発行から5秒以上経過したトークンの場合 401 を返す例
 *    if (isAuthorizationJwtExpired(req.headers, 1000 * 5)) return res(ctx.status(401));
 * })
 * ```
 */
export const isAuthorizationJwtExpired = (headers: Headers, tokenLifeMilliSec: number): boolean => {
  const value = headers.get("Authorization");
  if (!value) {
    return false;
  }
  if (!value.startsWith("Bearer eyJ")) {
    console.warn("[isAuthorizationJwtExpired] Invalid Bearer JWT format");
    return false;
  }
  const jwt = value.substring("Bearer ".length);
  const { iat } = decodeJwtPayload(jwt) as { iat: number };
  return Date.now() - iat > tokenLifeMilliSec;
};
