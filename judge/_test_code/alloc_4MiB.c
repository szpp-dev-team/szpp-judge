#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main() {
  size_t const len = 1024 * 1024 * 4;

  char *const a = malloc(len);
  if (a == NULL) {
    fprintf(stderr, "Failed to alloc %zu bytes\n", len);
    return 1;
  }

  // malloc() しただけでは物理メモリの割り当ては行われないので乱数を書き込む
  srand((unsigned)time(NULL));
  for (size_t i = 0; i < len; ++i) {
    a[i] = rand() % 128;
  }

  // 書き込むだけでは最適化で省略される可能性があるのでランダムに場所を選んで表示
  for (int i = 0; i < 5; ++i) {
    int j = rand() % len;
    printf("a[%d] = %d\n", j, a[j]);
  }
  printf("a[%d] = %d\n", 0, a[0]);
  printf("a[%zu] = %d\n", len - 1, a[len - 1]);

  return 0;
}
