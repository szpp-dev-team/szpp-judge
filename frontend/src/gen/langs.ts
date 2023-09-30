// @generated by github.com/szpp-dev-team/szpp-judge/langs/cmd/gen
// @generated on Sat, 30 Sep 2023 22:29:48 +0000

export const langIDs = [
  "c/11/gcc",
  "cpp/20/gcc",
  "java/21/openjdk",
  "python/3.11/cpython"
] as const;

export type LangID = (typeof langIDs)[number];

export interface LangMetaBrief {
  name: string;
  active: boolean;
}

export const langMetasBrief = {
  "c/11/gcc": {
    "name": "C11 (GCC 13.2)",
    "active": true
  },
  "cpp/20/gcc": {
    "name": "C++20 (GCC 13.2)",
    "active": true
  },
  "java/21/openjdk": {
    "name": "Java (OpenJDK 21)",
    "active": true
  },
  "python/3.11/cpython": {
    "name": "Python (CPython 3.11)",
    "active": true
  }
} as const;

export interface LangMetaFull {
  name: string;
  active: boolean;
  sourceFile: string;
  compileCmd: string[];
  execCmd: string[];
}

export const langMetasFull = {
  "c/11/gcc": {
    "name": "C11 (GCC 13.2)",
    "active": true,
    "sourceFile": "main.c",
    "compileCmd": [
      "gcc",
      "-std=c11",
      "-I/opt/include",
      "-lm",
      "-Wall",
      "-Wextra",
      "-DSZPP_JUDGE",
      "-O2",
      "-march=native",
      "-mtune=native",
      "main.c"
    ],
    "execCmd": [
      "./a.out"
    ]
  },
  "cpp/20/gcc": {
    "name": "C++20 (GCC 13.2)",
    "active": true,
    "sourceFile": "main.cpp",
    "compileCmd": [
      "g++",
      "-std=c++20",
      "-I/opt/include",
      "-lm",
      "-Wall",
      "-Wextra",
      "-DSZPP_JUDGE",
      "-O2",
      "-march=native",
      "-mtune=native",
      "main.cpp"
    ],
    "execCmd": [
      "./a.out"
    ]
  },
  "java/21/openjdk": {
    "name": "Java (OpenJDK 21)",
    "active": true,
    "sourceFile": "Main.java",
    "compileCmd": [
      "javac",
      "Main.java"
    ],
    "execCmd": [
      "java",
      "-Xss1G",
      "-Xmx1G",
      "Main"
    ]
  },
  "python/3.11/cpython": {
    "name": "Python (CPython 3.11)",
    "active": true,
    "sourceFile": "main.py",
    "compileCmd": [
      "python3",
      "-m",
      "compileall",
      "-q",
      "main.py"
    ],
    "execCmd": [
      "python3",
      "main.py"
    ]
  }
} as const;
