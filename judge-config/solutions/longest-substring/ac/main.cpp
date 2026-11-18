// longest-substring - Accepted solution (C++). Expected verdict: AC.

#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);  // an empty line (or EOF) means the empty string -> 0
    while (!s.empty() && (s.back() == '\r' || s.back() == '\n')) s.pop_back();

    // Sliding window; last[c] = most recent index of byte c (spaces count too).
    int last[256];
    fill(begin(last), end(last), -1);
    int best = 0, left = 0;
    for (int right = 0; right < (int)s.size(); ++right) {
        unsigned char c = s[right];
        if (last[c] >= left) left = last[c] + 1;
        last[c] = right;
        best = max(best, right - left + 1);
    }

    cout << best << "\n";
    return 0;
}
