import { LoginResponse } from "@/src/gen/proto/backend/v1/messages_pb";
import { AuthService } from "@/src/gen/proto/backend/v1/services-AuthService_connectquery";
import type { AuthUser } from "@/src/model/user";
import { Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

const REFRESH_TOKEN_PREFIX = "refreeeeesh";

type Credential = Pick<LoginResponse, "accessToken" | "refreshToken">;

export const generateMockAccessToken = (user: AuthUser, now: Date): string => {
  const payloadObj = {
    ...user,
    sub: "1234567890",
    iat: now.getTime(),
  };
  const header = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"; // { alg: "HS256", typ: "JWT" } のエンコード文字列
  const payload = Buffer.from(JSON.stringify(payloadObj)).toString("base64");
  const hash = "hassssh";
  return `${header}.${payload}.${hash}`;
};

export const generateMockRefreshToken = (now: Date): string => {
  // sv-SE ローケルでのフォーマットは `YYYY-mm-dd HH:MM:SS`
  return `${REFRESH_TOKEN_PREFIX}/${now.toLocaleString("sv-SE")}`;
};

export const generateMockCredential = (user: AuthUser): Credential => {
  const now = new Date();
  const accessToken = generateMockAccessToken(user, now);
  const refreshToken = generateMockRefreshToken(now);
  return { accessToken, refreshToken };
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

    const user: AuthUser = {
      id: 1,
      username,
      isAdmin: username === "admin",
    };

    return res(
      ctx.delay(500),
      encodeResp({
        user: {
          ...user,
          createdAt: Timestamp.now(),
        },
        ...generateMockCredential(user),
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
