import { LoginRequest, LogoutRequest } from "@/src/gen/proto/backend/v1/messages_pb";
import { AuthService } from "@/src/gen/proto/backend/v1/services_connect";
import { createPromiseClient, PromiseClient, Transport } from "@bufbuild/connect";
import { backendGrpcTransport } from "../config/grpc";
import type { AuthUser } from "../model/user";

export class AuthRepository {
  readonly #cli: PromiseClient<typeof AuthService>;

  constructor(t: Transport) {
    this.#cli = createPromiseClient(AuthService, t);
  }

  async login(req: LoginRequest): Promise<AuthUser> {
    const { user } = await this.#cli.login(req);
    return {
      ...user!,
    };
  }

  async logout(req: LogoutRequest): Promise<void> {
    await this.#cli.logout(req);
  }
}

export const authRepo = new AuthRepository(backendGrpcTransport);
