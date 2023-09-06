#include <stdio.h>

int main() {
  for (;;) {
    fputc('o', stdout);
    fputc('e', stderr);
  }
  return 0;
}
