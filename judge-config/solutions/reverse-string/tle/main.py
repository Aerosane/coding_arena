# reverse-string - slow solution (Python): repeated front insertion, O(N^2).
# Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import sys


def main():
    s = sys.stdin.readline().rstrip("\r\n")
    r = ""
    for c in s:
        r = c + r  # copies everything built so far
    print(r)


main()
