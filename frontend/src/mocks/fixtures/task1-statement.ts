const BQ = "`";

export const DUMMY_TASK1_STATEMENT = String.raw`
## 問題文

$A + B$ の答えを求めてください。

## 制約

- $1 \leq A, B \leq 10^{18}$

## 入力

入力は以下の形式で標準入力から与えられます。

${"```"}
$A$ $B$
$X_1$ $X_2$ ... $X_N$

$A_{1, 1}$ $\ldots$ $A_{1, W}$
$A_{2, 1}$ $\ldots$ $A_{2, W}$
$\vdots$
$A_{H, 1}$ ... $A_{H, W}$
${"```"}

## 出力

答えを出力してください。

## 出力
察してください。

## h2 ほげほげ

### h3 ほげほげ1
fuga~~~~~~~~~~~~~~~~
fuga~~~~~~~~~~~~~~~~

### h3 ほげほげ2
nya~~~~~~~~~~~~~~~~
nya~~~~~~~~~~~~~~~~

## 書式テスト

http://example.com

| left | center | right |
|:-----|:------:|------:|
| Bob  |   334  | 334   |
| Alice |   5   | 5     |

${"```"}cpp
#include <iostream>
#include <stdio.h>
int main() {
    printf("%d\n", 1 & 5);
    return 0;
    // </iostream>
}
${"```"}

**Double Asterisk**

__Double Undeerbar__

*Single Asterisk*

*Single Underbar*

~Single Tilde~

~~Double Tilde~~

## nested ol (indent=2)

1. aaaaa
2. bbbb
3. cccc
  1. XXXX
  2. YYYY
  3. CCCC
4. dddd

## nested ol (indent=4)

1. aaaaa
2. bbbb
3. cccc
    1. XXXX
    2. YYYY
    3. CCCC
4. dddd


## nested ul
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

---

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

<span>Here is ${BQ}span${BQ}</span>

<div>Here is ${BQ}div${BQ}</div>

<span class="example-class">Here is ${BQ}span${BQ} with ${BQ}class${BQ} attr</div>

<div class="example-class">Here is ${BQ}div${BQ} with ${BQ}class${BQ} attr</div>

↓ ${BQ}script${BQ} Tag
<script>
console.warn("XSS!! Hahahaha!!!!!!!!!!!");
</script>
↑ ${BQ}script${BQ} Tag

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
- $\text{\\${BQ}{a} \={a} \"{a} \v{a} \^{a} \u{a} \r{a}}$

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
`;
