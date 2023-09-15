import { atom, useAtomValue, useSetAtom } from "jotai";

interface ICredential {
  accessToken: string;
  refreshToken: string;
}

const credential = atom<ICredential>({ accessToken: "", refreshToken: "" });

export const useCredentialValue = () => useAtomValue(credential);
export const useCredentialSetter = () => useSetAtom(credential);
