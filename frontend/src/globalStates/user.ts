import { AuthUser } from "@/src/model/user";
import { atom, useAtomValue, useSetAtom } from "jotai";

const user = atom<AuthUser | null>(null);

export const useUserValue = () => useAtomValue(user);
export const useUserSetter = () => useSetAtom(user);
