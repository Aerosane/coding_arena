# longest-substring - Accepted solution (Python). Expected verdict: AC.

import sys


def main():
    s = sys.stdin.readline().rstrip("\r\n")  # empty line / EOF -> empty string -> 0

    # Sliding window; last[c] = most recent index of character c (spaces count too).
    last = {}
    best = 0
    left = 0
    for right, c in enumerate(s):
        if last.get(c, -1) >= left:
            left = last[c] + 1
        last[c] = right
        best = max(best, right - left + 1)

    print(best)


main()
