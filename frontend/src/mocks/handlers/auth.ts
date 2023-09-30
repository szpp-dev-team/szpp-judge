import { AuthService } from "@/src/gen/proto/backend/v1/auth_service-AuthService_connectquery";
import type { LoginResponse } from "@/src/gen/proto/backend/v1/auth_service_pb";
import type { AccessTokenClaim } from "@/src/model/auth";
import { Duration } from "@/src/util/time";
import { Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

const REFRESH_TOKEN_PREFIX = "refreeeeesh";

type Credential = Pick<LoginResponse, "accessToken" | "refreshToken">;

const generateMockAccessToken = (username: string, now: Date): string => {
  const payloadObj: AccessTokenClaim = {
    username,
    isAdmin: username.startsWith("admin"),
    iat: now.getTime(),
    exp: now.getTime() + Duration.SECOND * 10,
  };
  const header = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"; // { alg: "HS256", typ: "JWT" } のエンコード文字列
  const payload = Buffer.from(JSON.stringify(payloadObj)).toString("base64");
  const hash = "hassssh" + now.toLocaleTimeString();
  return `${header}.${payload}.${hash}`;
};

const generateMockRefreshToken = (username: string, now: Date): string => {
  // sv-SE ローケルでのフォーマットは `YYYY-mm-dd HH:MM:SS`
  return `${REFRESH_TOKEN_PREFIX}/${username}/${now.toLocaleString("sv-SE")}`;
};

const generateMockCredential = (username: string): Credential => {
  const now = new Date();
  const accessToken = generateMockAccessToken(username, now);
  const refreshToken = generateMockRefreshToken(username, now);
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

  grpcMock(AuthService, "refreshAccessToken", async (ctx, res, decodeReq, encodeResp) => {
    const { refreshToken } = await decodeReq();

    // 実際のリフレッシュトークンはスラッシュ区切りではない。
    // モックでは username の情報を埋め込むためにスラッシュ区切りにしている。
    const [prefix, username] = refreshToken.split("/");
    if (prefix !== REFRESH_TOKEN_PREFIX || !username) {
      return res(
        ctx.delay(100),
        ctx.status(401),
      );
    }

    const accessToken = generateMockAccessToken(username, new Date());
    return res(
      ctx.delay(100),
      encodeResp({ accessToken }),
    );
  }),
];
