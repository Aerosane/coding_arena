// valid-parentheses - Accepted solution (C++). Expected verdict: AC.

#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);  // an empty line (or EOF) means the empty string -> valid

    string stack;
    bool ok = true;
    for (char c : s) {
        if (c == '(' || c == '[' || c == '{') {
            stack.push_back(c);
        } else if (c == ')' || c == ']' || c == '}') {
            char open = (c == ')') ? '(' : (c == ']') ? '[' : '{';
            if (stack.empty() || stack.back() != open) {
                ok = false;
                break;
            }
            stack.pop_back();
        }
    }
    if (!stack.empty()) ok = false;

    cout << (ok ? "true" : "false") << "\n";
    return 0;
}
