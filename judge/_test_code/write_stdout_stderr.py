import sys
import time


n = int(input())
o = input()
e = input()

for _ in range(n):
    print(o, file=sys.stdout, flush=True)
    print(e, file=sys.stderr, flush=True)
    time.sleep(0.001)
