import { Testcase } from "@/src/gen/proto/backend/v1/task_resources_pb";
import { type PlainMessage } from "@bufbuild/protobuf";

export const dummySampleCases: Readonly<PlainMessage<Testcase>[]> = [
  {
    id: 1,
    slug: "sample_1",
    taskId: 1,
    input: "1 2 3",
    output: "Yes",
    description: String.raw`
この例では、すずっぴー君は $10^{18}$ 円を払うことで $20$ 円のおつりを貰います。

- hoge
- foo
- bar
`,
  },
  {
    id: 2,
    slug: "sample_2",
    taskId: 1,
    input: "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30",
    output: "Yeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeees",
    description: String.raw`横長テストです。`,
  },
  {
    id: 3,
    slug: "sample_2",
    taskId: 1,
    input: `
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20`,
    output: `Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
^^
Yes
Yes
Yes
No
`,
    description: String.raw`縦長テストです。`,
  },
] as const;
