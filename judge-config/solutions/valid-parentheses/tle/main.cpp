// valid-parentheses - slow solution (C++): repeated pair removal, O(N^2) on deeply nested input.
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);
    while (!s.empty() && (s.back() == '\r' || s.back() == '\n')) s.pop_back();

    // Keep deleting one adjacent matching pair per scan until nothing changes.
    const char *pairs[] = {"()", "[]", "{}"};
    bool changed = true;
    while (changed) {
        changed = false;
        for (const char *p : pairs) {
            size_t pos = s.find(p);
            if (pos != string::npos) {
                s.erase(pos, 2);
                changed = true;
            }
        }
    }

    cout << (s.empty() ? "true" : "false") << "\n";
    return 0;
}
