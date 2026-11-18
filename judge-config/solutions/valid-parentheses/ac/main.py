# valid-parentheses - Accepted solution (Python). Expected verdict: AC.

import sys


def main():
    s = sys.stdin.readline().strip()  # empty line / EOF -> empty string -> valid
    pairs = {")": "(", "]": "[", "}": "{"}
    stack = []
    ok = True
    for c in s:
        if c in "([{":
            stack.append(c)
        elif c in pairs:
            if not stack or stack[-1] != pairs[c]:
                ok = False
                break
            stack.pop()
    if stack:
        ok = False

    print("true" if ok else "false")


main()
