import type { AuthUser } from "@/src/model/user";
import { decodeJwtPayload } from "@/src/util/jwt";
import { atom, useAtom, useAtomValue, useSetAtom } from "jotai";
import { atomWithStorage, RESET } from "jotai/utils";

interface ICredential {
  accessToken: string;
  refreshToken: string;
}

const credential = atomWithStorage<ICredential>("szp_tkn", { accessToken: "", refreshToken: "" });

const authUser = atom<Readonly<AuthUser> | null>((get) => {
  const jwt = get(credential).accessToken;
  if (!jwt) {
    return null;
  }
  try {
    return decodeJwtPayload(jwt) as AuthUser;
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
