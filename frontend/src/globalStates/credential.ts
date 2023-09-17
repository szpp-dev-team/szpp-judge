import { useAtom, useAtomValue, useSetAtom } from "jotai";
import { atomWithStorage, RESET } from "jotai/utils";

interface ICredential {
  accessToken: string;
  refreshToken: string;
}

const credential = atomWithStorage<ICredential>("szp_tkn", { accessToken: "", refreshToken: "" });

export const useCredentialValue = () => useAtomValue(credential);
export const useCredentialSetter = () => useSetAtom(credential);

export const useCredentialValueAndEraser = (): [ICredential, () => void] => {
  const [cred, setCred] = useAtom(credential);
  return [cred, () => setCred(RESET)];
};
