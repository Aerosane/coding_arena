# two-sum - Accepted solution (Python). Expected verdict: AC.

import sys


def main():
    data = sys.stdin.read().split()
    n = int(data[0])
    nums = list(map(int, data[1:1 + n]))
    target = int(data[1 + n])

    # One pass with a hash map: value -> first index it appeared at.
    seen = {}
    for j, x in enumerate(nums):
        i = seen.get(target - x)
        if i is not None:
            print(i, j)
            return
        seen.setdefault(x, j)


main()
