#include <stdio.h>

int main() {
  for (;;) {
    fputc('o', stdout);
    fputc('E', stderr);
  }
  return 0;
}
