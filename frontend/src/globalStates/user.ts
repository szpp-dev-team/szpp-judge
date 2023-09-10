import { AuthUser } from "@/src/model/user";
import { atom, useAtom } from "jotai";

const user = atom<AuthUser | null>(null);

export const useUser = () => useAtom(user);
