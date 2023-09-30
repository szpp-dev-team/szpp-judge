import { decodeJwtPayload } from "@/src/util/jwt";

/**
 * Authorization ヘッダの値を JWT としてパースし、
 * トークンの発行時刻と `tokenLifeMilliSec` から期限切れか否かを判定する。
 *
 * 使用例:
 * ```ts
 * grpcMock(..., async (ctx, res, decodeReq, encodeResp, req) => {
 *    if (isAuthorizationJwtExpired(req.headers)) return res(ctx.status(401));
 * })
 * ```
 */
export const isAuthorizationJwtExpired = (headers: Headers, tokenLifeMilliSec: number): boolean => {
  const jwt = headers.get("Authorization");
  if (!jwt) {
    return false;
  }
  const { iat } = decodeJwtPayload(jwt) as { iat: number };
  return Date.now() - iat > tokenLifeMilliSec;
};
