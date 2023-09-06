#include <stdio.h>
#include <string.h>

int main() {
  unsigned N, K;
  if (scanf("%u%u", &N, &K) != 2) {
    fprintf(stderr, "Failed to read two integers from stdin\n");
    return 1;
  }

  if (K > 100) {
    fprintf(stderr, "Too many file num specification: %u\n", K);
    return 1;
  }

  // 1 MiB のブロック。この内容を1ファイルあたり n 回書き込む。
  char block[1024 * 1024];
  memset(block, 'x', sizeof(block));

  for (unsigned i = 0; i < K; ++i) {
    char filename[16];
    sprintf(filename, "%u.txt", i + 1);

    FILE *fp = fopen(filename, "w");
    if (fp == NULL) {
      fprintf(stderr, "Failed to create a file '%s'\n", filename);
      return 1;
    }

    for (unsigned j = 0; j < N; ++j) {
      fwrite(block, sizeof(block), 1, fp);
    }
    fclose(fp);
  }
}
