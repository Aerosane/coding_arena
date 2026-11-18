# two-sum - slow solution (Python): brute force over all pairs, O(N^2).
# Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import sys


def main():
    data = sys.stdin.read().split()
    n = int(data[0])
    a = list(map(int, data[1:1 + n]))
    target = int(data[1 + n])

    for i in range(n):
        for j in range(i + 1, n):
            if a[i] + a[j] == target:
                print(i, j)
                return


main()
