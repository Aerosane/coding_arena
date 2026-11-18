# merge-sorted-arrays - slow solution (Python): concatenate then selection sort, O(K^2).
# Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import sys


def main():
    data = sys.stdin.read().split()
    m, n = int(data[0]), int(data[1])
    v = list(map(int, data[2:2 + m + n]))  # nums1 then nums2

    k = len(v)
    for i in range(k):  # selection sort
        mn = i
        for j in range(i + 1, k):
            if v[j] < v[mn]:
                mn = j
        v[i], v[mn] = v[mn], v[i]

    print(" ".join(map(str, v)))


main()
