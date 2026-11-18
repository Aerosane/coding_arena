# Reference solutions for judge testing

Layout: `<problem>/<ac|tle>/{main.cpp, main.py, Main.java}` (Java class is `Main`).

| Problem | AC algorithm | TLE algorithm |
|---|---|---|
| two-sum | one-pass hash map, O(N) | brute force over all pairs, O(N^2) |
| reverse-string | reverse, O(N) | repeated front insertion, O(N^2) |
| valid-parentheses | stack, O(N) | repeated pair removal, O(N^2) on nested input |
| merge-sorted-arrays | two-pointer merge, O(M+N) | concatenate + selection sort, O(K^2) |
| longest-substring | sliding window, O(N) | scan every start to the end, no early break, O(N^2) |

## Verification

    python judge-config/solutions/run_solutions.py            # official tests + stress
    python judge-config/solutions/run_solutions.py --no-stress

* Phase 1 (official tests): every solution, AC and TLE, must be correct.
* Phase 2 (generated stress inputs): AC must still pass, TLE must time out.
  Several stress inputs exceed the official constraints on purpose (two-sum
  N=2e5, reverse-string 1e6, parentheses 4e5, merge m+n=2e5, substring 2e5).
  They are generated in memory and are NOT added to `judge-config/problems/`.

## Known limitations (read before expecting TLE verdicts on the judge)

The official tests are tiny (largest file about 1.6 KB), so every TLE solution is
**Accepted** on the current judge data. Even at the official *maximum* sizes,
measured locally against the 1.0 s limit:

| Problem (TLE solution, max official size) | C++ | Python | Java |
|---|---|---|---|
| two-sum (N=1e4) | 0.02s AC | 2.17s TLE | 0.13s AC |
| reverse-string (1e5) | 0.10s AC | 0.17s AC | 0.23s AC |
| valid-parentheses (1e4) | 0.09s AC | 0.08s AC | 0.11s AC |
| merge-sorted-arrays (m+n=200) | 0.00s AC | 0.01s AC | 0.05s AC |
| longest-substring (5e4) | 0.25s AC | 37.1s TLE | 1.17s borderline |

To get real TLE verdicts for the others, the judge needs larger test cases and
higher constraints (merge-sorted-arrays with m+n <= 200 can never TLE).

`reverse-string`: `problems.json` describes `n` + space-separated chars, but the
real `.in` files are one raw line; solutions follow the real test data.
Timings are from a Linux sandbox, not from the DMOJ judge.
