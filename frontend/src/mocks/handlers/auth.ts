import { LoginRequest, LoginResponse } from "@/src/gen/proto/backend/v1/messages_pb";
import { AuthService } from "@/src/gen/proto/backend/v1/services_connect";
import { Timestamp } from "@bufbuild/protobuf";
import { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

export const authHandlers: RequestHandler[] = [
  grpcMock(AuthService, "login", async (ctx, res, decodeReq, encodeResp) => {
    const { username, password } = await decodeReq() as LoginRequest;
    const ok = (username === "user" || username === "admin") && password === "pass";
    if (!ok) {
      return res(
        ctx.delay(500),
        ctx.status(401), // 401 Forbidden
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
        }),
      ),
    );
  }),
];
