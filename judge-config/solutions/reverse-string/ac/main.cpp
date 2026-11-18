// reverse-string - Accepted solution (C++). Expected verdict: AC.

#include <bits/stdc++.h>
using namespace std;

int main() {
    // Test data is one raw line (spaces are significant), e.g. "hello world".
    string s;
    getline(cin, s);
    while (!s.empty() && (s.back() == '\r' || s.back() == '\n')) s.pop_back();

    reverse(s.begin(), s.end());

    cout << s << "\n";
    return 0;
}
