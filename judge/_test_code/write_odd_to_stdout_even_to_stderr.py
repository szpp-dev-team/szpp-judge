import sys


n = int(input())

for i in range(n):
    if i & 1:
        print(i, file=sys.stdout)
    else:
        print(i, file=sys.stderr)
