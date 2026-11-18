// two-sum - Accepted solution (C++). Expected verdict: AC.

#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    if (!(cin >> n)) return 0;
    vector<long long> a(n);
    for (auto &x : a) cin >> x;
    long long target;
    cin >> target;

    // One pass with a hash map: value -> first index it appeared at.
    unordered_map<long long, int> seen;
    seen.reserve(2 * (size_t)n);
    for (int j = 0; j < n; ++j) {
        auto it = seen.find(target - a[j]);
        if (it != seen.end()) {
            cout << it->second << " " << j << "\n";
            return 0;
        }
        seen.emplace(a[j], j);
    }
    return 0;
}
