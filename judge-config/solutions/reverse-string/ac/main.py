# reverse-string - Accepted solution (Python). Expected verdict: AC.

import sys


def main():
    # Test data is one raw line (spaces are significant), e.g. "hello world".
    s = sys.stdin.readline().rstrip("\r\n")
    result = s[::-1]

    print(result)


main()
