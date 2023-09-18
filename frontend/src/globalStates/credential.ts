import { decodeJwtPayload } from "@/src/util/jwt";
import { atom, useAtom, useAtomValue, useSetAtom } from "jotai";
import { atomWithStorage, RESET } from "jotai/utils";

export interface ICredential {
  accessToken: string;
  refreshToken: string;
}

export type IAuthUser = {
  username: string;
  isAdmin: boolean;
};

const credential = atomWithStorage<ICredential>("szp_tkn", { accessToken: "", refreshToken: "" });

const authUser = atom<Readonly<IAuthUser> | null>((get) => {
  const jwt = get(credential).accessToken;
  if (!jwt) {
    return null;
  }
  try {
    return decodeJwtPayload(jwt) as IAuthUser;
  } catch (e) {
    console.error(e);
    return null;
  }
});

export const useCredentialValue = () => useAtomValue(credential);
export const useCredentialSetter = () => useSetAtom(credential);

export const useCredentialValueAndEraser = (): [ICredential, () => void] => {
  const [cred, setCred] = useAtom(credential);
  return [cred, () => setCred(RESET)];
};

export const useAuthUserValue = () => useAtomValue(authUser);
