#include <stddef.h>
#include <stdio.h>
#include <string.h>

int main() {
  size_t n;
  if (scanf("%zu", &n) != 1) {
    fprintf(stderr, "Failed to read an integer from stdin\n");
    return 1;
  }

  FILE *fp = fopen("x.txt", "w");
  if (fp == NULL) {
    fprintf(stderr, "Failed to create a file\n");
    return 1;
  }

  char buf[1024 * 1024];
  memset(buf, 'x', sizeof(buf));

  while (n--) {
    fwrite(buf, sizeof(buf), 1, fp);
  }

  fclose(fp);
}
