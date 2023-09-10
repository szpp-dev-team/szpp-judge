import { useCredentialSetter } from "@/src/globalStates/credential";
import { useUserSetter } from "@/src/globalStates/user";
import { AuthUser } from "@/src/model/user";
import { authRepo, AuthRepository } from "@/src/repository/auth";

const createAuthUsecase = ({ repository, setCredential, setUser }: {
  repository: AuthRepository;
  setCredential: ReturnType<typeof useCredentialSetter>;
  setUser: ReturnType<typeof useUserSetter>;
}) => {
  return {
    // rest parameter を使うとこの関数で使いたい repository.<method> が1つなら
    // いいけどそれ以上のことをやらせたいときに引数が自由に使えなくて不便かも
    // ただできればこのレイヤーは protobuf のことを知らなくていいようにしておきたい
    async login(...params: Parameters<typeof repository.login>): Promise<AuthUser> {
      const resp = await repository.login(...params);

      // globalState 更新
      setCredential({ refreshToken: "tooooooooooooooooken" }); // TODO: ちゃんとトークンをセットする
      setUser(resp);

      return resp;
    },
    async logout(): Promise<void> {
    },
  };
};

export const useAuthUsecase = () => {
  const credentialSetter = useCredentialSetter();
  const userSetter = useUserSetter(); // 単に setUser でもいいかな
  return createAuthUsecase({ repository: authRepo, setCredential: credentialSetter, setUser: userSetter }); // TODO: React.memo
};
