import { LoginResponse } from "@/src/gen/proto/backend/v1/messages_pb";
import { AuthService } from "@/src/gen/proto/backend/v1/services-AuthService_connectquery";
import { Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

const ACCESS_TOKEN_PREFIX = "eyJxxx";
const REFRESH_TOKEN_PREFIX = "refreeeeesh";

type Credential = Pick<LoginResponse, "accessToken" | "refreshToken">;
type AccessTokenClaim = {
  username: string;
  isAdmin: boolean;
};

export const generateMockCredential = (username: string): Credential => {
  const now = new Date().toLocaleString("sv-SE"); // JST で "YYYY-mm-dd HH:MM:SS" を返してくれる
  return {
    accessToken: `${ACCESS_TOKEN_PREFIX}/${username}/${now}`,
    refreshToken: `${REFRESH_TOKEN_PREFIX}/${now}`,
  };
};

// パースに失敗した場合 or 不正なトークンの場合は null を返す (モックなので厳密なチェックはしない)
export const parseAccessToken = (accessToken: string): AccessTokenClaim | null => {
  const a = accessToken.split("/");
  const [prefix, username] = a;
  if (a.length !== 3 || prefix !== ACCESS_TOKEN_PREFIX) {
    return null;
  }
  const isAdmin = username === "admin";
  return { username, isAdmin };
};

export const authHandlers: RequestHandler[] = [
  grpcMock(AuthService, "login", async (ctx, res, decodeReq, encodeResp) => {
    const { username, password } = await decodeReq();
    const ok = (username === "user" || username === "admin") && password === "pass";
    if (!ok) {
      return res(
        ctx.delay(500),
        ctx.status(401), // 401 Unauthorized
      );
    }

    return res(
      ctx.delay(500),
      encodeResp({
        user: {
          id: 1,
          username,
          isAdmin: username === "admin",
          createdAt: Timestamp.now(),
        },
        ...generateMockCredential(username),
      }),
    );
  }),

  grpcMock(AuthService, "logout", async (ctx, res, decodeReq, encodeResp) => {
    const { refreshToken } = await decodeReq();

    const ok = refreshToken.startsWith(REFRESH_TOKEN_PREFIX);
    if (!ok) {
      return res(
        ctx.delay(500),
        ctx.status(401),
      );
    }
    return res(
      ctx.delay(500),
      encodeResp({}),
    );
  }),
];
