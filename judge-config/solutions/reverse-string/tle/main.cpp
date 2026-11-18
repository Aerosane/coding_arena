// reverse-string - slow solution (C++): repeated front insertion, O(N^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);
    while (!s.empty() && (s.back() == '\r' || s.back() == '\n')) s.pop_back();

    string r;
    for (char c : s) r.insert(r.begin(), c);  // shifts everything built so far

    cout << r << "\n";
    return 0;
}
