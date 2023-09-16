import { useAtomValue, useSetAtom } from "jotai";
import { atomWithStorage } from "jotai/utils";

interface ICredential {
  accessToken: string;
  refreshToken: string;
}

const credential = atomWithStorage<ICredential>("szp_tkn", { accessToken: "", refreshToken: "" });

export const useCredentialValue = () => useAtomValue(credential);
export const useCredentialSetter = () => useSetAtom(credential);
