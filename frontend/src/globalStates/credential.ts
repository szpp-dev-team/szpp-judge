import { atom, useAtomValue, useSetAtom } from "jotai";

interface ICredential {
  refreshToken: string;
}

const credential = atom<ICredential>({ refreshToken: "" });

export const useCredentialValue = () => useAtomValue(credential);
export const useCredentialSetter = () => useSetAtom(credential);
