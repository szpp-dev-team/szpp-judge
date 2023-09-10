import { AuthUser } from "@/src/model/user";
import { authRepo, AuthRepository } from "@/src/repository/auth";

const createAuthUsecase = ({ repository }: { repository: AuthRepository }) => {
  return {
    // rest parameter を使うとこの関数で使いたい repository.<method> が1つなら
    // いいけどそれ以上のことをやらせたいときに引数が自由に使えなくて不便かも
    async login(...args: Parameters<typeof repository.login>): Promise<AuthUser> {
      const resp = await repository.login(...args);
      return resp;
    },
    async logout(): Promise<void> {
    },
  };
};

export const useAuthUsecase = () => {
  return createAuthUsecase({ repository: authRepo }); // TODO: React.memo
};
