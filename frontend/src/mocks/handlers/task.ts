import { Difficulty, Task } from "@/src/gen/proto/backend/v1/task_resources_pb";
import type { PlainMessage } from "@bufbuild/protobuf";
import type { RequestHandler } from "msw";

export const dummyTasks: Readonly<PlainMessage<Task>[]> = [
  {
    id: 1000,
    title: "マークダウン・KaTeXテスト1",
    execTimeLimit: 1000,
    execMemoryLimit: 256,
    difficulty: Difficulty.BEGINNER,
    statement: String.raw`
## 問題文

$A + B$ の答えを求めてください。

## 制約

- $1 \leq A, B \leq 10^{18}$

## 入力

入力は以下の形式で標準入力から与えられます。

\`\`\`
$A$ $B$
$X_1$ $X_2$ ... $X_N$
\`\`\`

## 出力

答えを出力してください。


## 書式テスト

http://example.com

| left | center | right |
|:-----|:------:|------:|
| Bob  |   334  | 334   |
| Alice |   5   | 5     |

\`\`\`cpp
#include <iostream>
#include <stdio.h>
int main() {
    printf("%d\n", 1 & 5);
    return 0;
    // </iostream>
}
\`\`\`

- KaTeX
- $\KaTeX$

$$
\Gamma(z) = \int_0^\infty t^{z-1}e^{-t}dt\,.
$$

This is x: $x$

- micromark の README よりコピーした TeX

Lift($L$) can be determined by Lift Coefficient ($C_L$) like the following equation.

$$
L = \frac{1}{2} \rho v^2 S C_L
$$

## Raw HTML tags

<span>Here is \`span\`</span>

<div>Here is \`div\`</div>

<span class="example-class">Here is \`span\` with \`class\` attr</div>

<div class="example-class">Here is \`div\` with \`class\` attr</div>

↓ \`script\` Tag
<script>
console.warn("XSS!! Hahahaha!!!!!!!!!!!");
</script>
↑ \`script\` Tag

0 < 1 < 2

9 > 8 > 7

<details>
  <summary>Here is Summary</summary>
  Detail Detail Detail Detail
  Detail Detail Detail Detail
  Detail Detail Detail Detail
</details>

## Escape Characters
参考: https://katex.org/docs/supported.html#symbols-and-punctuation

- \# \$ \% \& \_ \{ \}
- $ \# \$ \% \& \_ \{ \} $

Escaped Dollar marker is \$ a + b \$ , Dollar in KaTeX is $ \$ a + b \$ $

$$
\sum_{i=0}^{n} \left(\log \mathrm{sim}(Q^{(i)}, k_+) \right)^2 $a
$$

## Accent functions

- \' \= \" \v \^ \u \r
- $\text{\\\`{a} \={a} \"{a} \v{a} \^{a} \u{a} \r{a}}$

## Spaces

参考: https://katex.org/docs/supported.html#overlap-and-spacing

- a\,a\>a\:a\;a~a
- $a\,a\>a\:a\;a~a$

## Environments
参考: https://katex.org/docs/supported.html#environments

$$
x = \begin{cases}
   a &\text{if } b \\
   c &\text{if } d
\end{cases}
$$

---

$$
\def\arraystretch{1.5}
   \begin{array}{c:c:c}
   a & b & c \\ \hline
   d & e & f \\
   \hdashline
   g & h & i
\end{array}
$$

---

$$
\begin{CD}
   A @>a>> B   \\
@VbVV    @AAcA \\
   C   @=  D
\end{CD}
$$

---

inline katex which contains newline: $a
b
c
$

$A$$B$

block katex without outer space$$\frac{a}{b}$$xxxx
`,
  },
  {
    id: 1001,
    title: "マークダウン・KaTeXテスト2",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.EASY,
    statement: String.raw`
## 問題文
すずっぴー君はちらし寿司ですか？

**Double Asterisk**

__Double Undeerbar__

*Single Asterisk*

*Single Underbar*

~Single Tilde~

~~Double Tilde~~


## 制約
- aaaa
- bbbb
  - bbbb-1
  - bbbb-2
- cccc
  - cccc-1
    - cccc-1-1
    - cccc-1-2
    - cccc-1-3
  - cccc-2
  - cccc-3

## 入力
\`\`\`
$A_{1, 1}$ $\ldots$ $A_{1, W}$
$A_{2, 1}$ $\ldots$ $A_{2, W}$
$\vdots$
$A_{H, 1}$ ... $A_{H, W}$
\`\`\`

## 出力
察してください。

## h2 ほげほげ

### h3 ほげほげ1
fuga~~~~~~~~~~~~~~~~
fuga~~~~~~~~~~~~~~~~

### h3 ほげほげ2
nya~~~~~~~~~~~~~~~~
nya~~~~~~~~~~~~~~~~
`,
  },
  {
    id: 1002,
    title: "すずっぴー君のおつかい",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.EASY,
    statement: "## 問題文\nすずっぴー君はおつかいに行きました。 ",
  },
  {
    id: 1003,
    title: "すずっぴー君の怒り",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.EASY,
    statement: "## 問題文\nすずっぴー君は激怒した。 ",
  },
  {
    id: 1004,
    title: "すずっぴー君の分裂",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.MEDIUM,
    statement: "## 問題文\nすずっぴー君は分裂を始めました。次はあなたの番です。 ",
  },
  {
    id: 1005,
    title: "すずっぴー君の爆散",
    execTimeLimit: 1000,
    execMemoryLimit: 512,
    difficulty: Difficulty.MEDIUM,
    statement: "## 問題文\nバーン。 ",
  },
  {
    id: 1006,
    title: "師匠との出会い",
    execTimeLimit: 2000,
    execMemoryLimit: 1024,
    difficulty: Difficulty.HARD,
    statement: "## 問題文\nにゃーん。 ",
  },
  {
    id: 1007,
    title: "師匠の爆散",
    execTimeLimit: 5000,
    execMemoryLimit: 2048,
    difficulty: Difficulty.HARD,
    statement: "## 問題文\nバーンバーンバーン。 ",
  },
  {
    id: 1008,
    title:
      "長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル",
    execTimeLimit: 500,
    execMemoryLimit: 256,
    difficulty: Difficulty.IMPOSSIBLE,
    statement: "## 問題文\nすずっぴー君はおつかいに行きました。 ",
  },
] as const;

export const taskHandlers: RequestHandler[] = [];
