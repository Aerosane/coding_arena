# merge-sorted-arrays - Accepted solution (Python). Expected verdict: AC.

import sys


def main():
    # Token-based parsing: an array line may be empty when m or n is 0.
    data = sys.stdin.read().split()
    m, n = int(data[0]), int(data[1])
    a = list(map(int, data[2:2 + m]))
    b = list(map(int, data[2 + m:2 + m + n]))

    out = []
    i = j = 0
    while i < m and j < n:
        if a[i] <= b[j]:
            out.append(a[i])
            i += 1
        else:
            out.append(b[j])
            j += 1
    out.extend(a[i:])
    out.extend(b[j:])

    print(" ".join(map(str, out)))


main()
