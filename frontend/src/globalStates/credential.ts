import { atom, useAtom } from "jotai";

interface ICredential {
  accessToken: string;
  refreshToken: string;
}

const credential = atom<ICredential>({ accessToken: "", refreshToken: "" });

export const useCredential = () => useAtom(credential);
