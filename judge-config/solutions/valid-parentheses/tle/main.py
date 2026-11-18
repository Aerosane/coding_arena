# valid-parentheses - slow solution (Python): repeated pair removal, O(N^2) on deeply nested input.
# Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import sys


def main():
    s = sys.stdin.readline().strip()

    # Keep deleting adjacent matching pairs until nothing changes.
    # Nested input like "((()))" only loses one pair per pass.
    while True:
        t = s.replace("()", "").replace("[]", "").replace("{}", "")
        if t == s:
            break
        s = t

    print("true" if not s else "false")


main()
