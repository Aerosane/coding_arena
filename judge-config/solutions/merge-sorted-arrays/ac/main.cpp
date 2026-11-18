// merge-sorted-arrays - Accepted solution (C++). Expected verdict: AC.

#include <bits/stdc++.h>
using namespace std;

int main() {
    int m, n;
    if (!(cin >> m >> n)) return 0;
    vector<long long> a(m), b(n);
    for (auto &x : a) cin >> x;  // an array line may be empty when m or n is 0
    for (auto &x : b) cin >> x;

    // Classic two-pointer merge.
    vector<long long> out;
    out.reserve(m + n);
    int i = 0, j = 0;
    while (i < m && j < n) out.push_back(a[i] <= b[j] ? a[i++] : b[j++]);
    while (i < m) out.push_back(a[i++]);
    while (j < n) out.push_back(b[j++]);

    for (size_t k = 0; k < out.size(); ++k) cout << (k ? " " : "") << out[k];
    cout << "\n";
    return 0;
}
