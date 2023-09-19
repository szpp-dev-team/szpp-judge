import { AuthService } from "@/src/gen/proto/backend/v1/auth_service-AuthService_connectquery";
import { LoginRequest, LoginResponse } from "@/src/gen/proto/backend/v1/auth_service_pb";
import { Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

export const authHandlers: RequestHandler[] = [
  grpcMock(AuthService, "login", async (ctx, res, decodeReq, encodeResp) => {
    const { username, password } = await decodeReq() as LoginRequest;
    const ok = (username === "user" || username === "admin") && password === "pass";
    if (!ok) {
      return res(
        ctx.delay(500),
        ctx.status(401), // 401 Unauthorized
      );
    }

    return res(
      ctx.delay(500),
      encodeResp(
        new LoginResponse({
          user: {
            id: 1,
            username,
            isAdmin: username === "admin",
            createdAt: Timestamp.now(),
          },
          accessToken: "eyJ..." + Date.now(), // 適当
          refreshToken: "refreeeeesh" + Date.now(), // 適当
        }),
      ),
    );
  }),
];
