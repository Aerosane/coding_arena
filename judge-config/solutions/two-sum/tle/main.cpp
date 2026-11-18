// two-sum - slow solution (C++): brute force over all pairs, O(N^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    if (!(cin >> n)) return 0;
    vector<long long> a(n);
    for (auto &x : a) cin >> x;
    long long target;
    cin >> target;

    for (int i = 0; i < n; ++i) {
        for (int j = i + 1; j < n; ++j) {
            if (a[i] + a[j] == target) {
                cout << i << " " << j << "\n";
                return 0;
            }
        }
    }
    return 0;
}
