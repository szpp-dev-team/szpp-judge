import type { AuthUser, Credential } from "@/src/model/auth";
import { decodeJwtPayload } from "@/src/util/jwt";
import { atom, useAtom, useAtomValue, useSetAtom } from "jotai";
import { atomWithStorage, RESET } from "jotai/utils";

const credential = atomWithStorage<Readonly<Credential>>("szp_tkn", { accessToken: "", refreshToken: "" });

const authUser = atom<Readonly<AuthUser> | null>((get) => {
  const jwt = get(credential).accessToken;
  if (!jwt) {
    return null;
  }
  try {
    return decodeJwtPayload(jwt) as Readonly<AuthUser>;
  } catch (e) {
    console.error(e);
    return null;
  }
});

export const useCredentialValue = () => useAtomValue(credential);
export const useCredentialSetter = () => useSetAtom(credential);
export const useCredentialValueAndSetter = () => useAtom(credential);

export const useCredentialValueAndEraser = (): [Readonly<Credential>, () => void] => {
  const [cred, setCred] = useAtom(credential);
  return [cred, () => setCred(RESET)];
};

export const useAuthUserValue = () => useAtomValue(authUser);
