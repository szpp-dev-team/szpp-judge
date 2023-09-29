#include <stdio.h>

int main() {
  for (;;) {
    fputc('X', stderr);
  }
  return 0;
}
