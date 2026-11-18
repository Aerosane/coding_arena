// longest-substring - slow solution (C++): scan every start to the end with no early break, O(N^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);
    while (!s.empty() && (s.back() == '\r' || s.back() == '\n')) s.pop_back();

    int n = (int)s.size();
    int stamp[256] = {0};
    int best = 0;
    for (int i = 0; i < n; ++i) {
        // Scan from i to the end every time, even after a repeat was seen.
        bool dup = false;
        int len = 0;
        for (int j = i; j < n; ++j) {
            unsigned char c = s[j];
            if (!dup) {
                if (stamp[c] == i + 1) dup = true;
                else { stamp[c] = i + 1; ++len; }
            }
        }
        best = max(best, len);
    }

    cout << best << "\n";
    return 0;
}
