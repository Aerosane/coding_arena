# longest-substring - slow solution (Python): scan every start to the end with no early break, O(N^2).
# Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import sys


def main():
    s = sys.stdin.readline().rstrip("\r\n")
    n = len(s)
    best = 0
    for i in range(n):
        # Scan from i to the end every time, even after a repeat was seen.
        seen = set()
        dup = False
        length = 0
        for j in range(i, n):
            c = s[j]
            if not dup:
                if c in seen:
                    dup = True
                else:
                    seen.add(c)
                    length += 1
        best = max(best, length)
    print(best)


main()
