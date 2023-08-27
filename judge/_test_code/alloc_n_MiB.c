#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/resource.h>
#include <sys/types.h>
#include <time.h>
#include <unistd.h>

// stdin から整数 n を読み取って `n` MiB の領域を確保する。

void show_stats(const char *title) {
  char buf[64];

  printf("============ %s ===========\n", title);
  fflush(stdout);
  fflush(stderr);

  system("ps aux");
  fflush(stdout);

  pid_t const pid = getpid();
  sprintf(buf, "grep '^Vm' /proc/%d/status", pid);
  system(buf);

  struct rusage u;
  getrusage(RUSAGE_SELF, &u);
  printf("maxRSS from getrusage(): %lu [MiB]\n", u.ru_maxrss / 1024);
  fflush(stdout);
}

int main() {
  printf("Self pid = %d\n", getpid());
  show_stats("Initial");

  size_t n;
  if (scanf("%zu", &n) != 1) {
    fprintf(stderr, "Failed to read an integer from stdin");
    return 1;
  }

  size_t const len = 1024 * 1024 * n;

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

  show_stats("allocating");
  usleep(1000 * 500);

  free(a);

  show_stats("after free");

  return 0;
}
