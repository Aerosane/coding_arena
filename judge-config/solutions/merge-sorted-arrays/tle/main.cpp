// merge-sorted-arrays - slow solution (C++): concatenate then selection sort, O(K^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

#include <bits/stdc++.h>
using namespace std;

int main() {
    int m, n;
    if (!(cin >> m >> n)) return 0;
    vector<long long> v(m + n);
    for (auto &x : v) cin >> x;  // nums1 then nums2, ignoring that each is already sorted

    int k = m + n;
    for (int i = 0; i < k; ++i) {  // selection sort
        int mn = i;
        for (int j = i + 1; j < k; ++j)
            if (v[j] < v[mn]) mn = j;
        swap(v[i], v[mn]);
    }

    for (int i = 0; i < k; ++i) cout << (i ? " " : "") << v[i];
    cout << "\n";
    return 0;
}
