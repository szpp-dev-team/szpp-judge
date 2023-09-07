import { CreateUserRequest } from "@/src/gen/proto/backend/v1/messages_pb";
import { UserService } from "@/src/gen/proto/backend/v1/services_connect";
import { createPromiseClient, PromiseClient, Transport } from "@bufbuild/connect";
import { backendGrpcTransport } from "../config/grpc";
import type { AuthUser } from "../model/user";

export class UserRepository {
  readonly cli: PromiseClient<typeof UserService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(UserService, t);
  }

  async register(req: CreateUserRequest): Promise<AuthUser> {
    const { user } = await this.cli.createUser(req);
    return {
      ...user!,
    };
  }
}

export const userRepo = new UserRepository(backendGrpcTransport);
