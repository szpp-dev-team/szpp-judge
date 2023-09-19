import { UserService } from "@/src/gen/proto/backend/v1/user_service-UserService_connectquery";
import { CreateUserRequest, CreateUserResponse } from "@/src/gen/proto/backend/v1/user_service_pb";
import { Timestamp } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";
import { grpcMock } from "../grpc";

export const userHandlers: RequestHandler[] = [
  grpcMock(UserService, "createUser", async (ctx, res, decodeReq, encodeResp) => {
    const { username, password, email } = await decodeReq() as CreateUserRequest;
    const ok = (username === "user" || username === "admin") && password === "Pass.w0rd" && email === "user@hoge.test";
    if (!ok) {
      return res(
        ctx.delay(500),
        ctx.status(400), // 400 Bad Request
      );
    }
    return res(
      ctx.delay(500),
      encodeResp(
        new CreateUserResponse({
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
