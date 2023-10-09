import { type LangID } from "../gen/langs";

export const MAX_SOURCE_CODE_SIZE = 512 * 1024; // 512 KiB

export const acceptedFileExts = [
  ".c",
  ".cpp",
  ".cc",
  ".py",
  ".java",
  ".sb3",
] as const;

export type AcceptedFileExt = (typeof acceptedFileExts)[number];

// NOTE* 将来、提出可能な言語が増えて C++17 や C++20 の clang 版などができた場合に対応・仕様の見直しが必要
export const fileExt2langId = {
  ".c": "c/11/gcc",
  ".cc": "cpp/20/gcc",
  ".cpp": "cpp/20/gcc",
  ".py": "python/3.11/cpython",
  ".java": "java/21/openjdk",
  ".sb3": "scratch/3/gcc",
} satisfies Record<AcceptedFileExt, LangID>;
