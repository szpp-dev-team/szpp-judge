import { useCredential } from "@/src/globalStates/credential";
import { useUser } from "@/src/globalStates/user";
import { AuthUser } from "@/src/model/user";
import { authRepo, AuthRepository } from "@/src/repository/auth";

const createAuthUsecase = ({ repository, setCredential, setUser }: { repository: AuthRepository, setCredential: ReturnType<typeof useCredential>[1], setUser: ReturnType<typeof useUser>[1] }) => {
  return {
    // rest parameter を使うとこの関数で使いたい repository.<method> が1つなら
    // いいけどそれ以上のことをやらせたいときに引数が自由に使えなくて不便かも
    async login(...args: Parameters<typeof repository.login>): Promise<AuthUser> {
      const resp = await repository.login(...args);

      // globalState 更新
      setCredential({ refreshToken: "tooooooooooooooooken" }) // TODO: ちゃんとトークンをセットする
      setUser(resp)

      return resp;
    },
    async logout(): Promise<void> {
    },
  };
};

export const useAuthUsecase = () => {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [_c, setCredential] = useCredential();
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [_u, setUser] = useUser();
  return createAuthUsecase({ repository: authRepo, setCredential, setUser }); // TODO: React.memo
};
