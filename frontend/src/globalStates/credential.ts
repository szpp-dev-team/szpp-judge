import { atom, useAtom } from "jotai";

interface ICredential {
  refreshToken: string;
}

const credential = atom<ICredential>({ refreshToken: "" });

export const useCredential = () => useAtom(credential);
