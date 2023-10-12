import { useAtomValue, useSetAtom } from "jotai";
import { atomWithStorage } from "jotai/utils";
import { LangID } from "../gen/langs";

const langId = atomWithStorage<Readonly<LangID>>("szp_submit_lang_id", "cpp/20/gcc");

export const useSubmissionLangValue = () => useAtomValue(langId);
export const useSubmissionLangSetter = () => useSetAtom(langId);
