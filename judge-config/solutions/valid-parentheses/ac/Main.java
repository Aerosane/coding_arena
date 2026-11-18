// valid-parentheses - Accepted solution (Java). Expected verdict: AC.

import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String s = br.readLine();
        if (s == null) s = "";  // empty line / EOF -> empty string -> valid

        char[] stack = new char[s.length()];
        int top = 0;
        boolean ok = true;
        for (int k = 0; k < s.length() && ok; k++) {
            char c = s.charAt(k);
            if (c == '(' || c == '[' || c == '{') {
                stack[top++] = c;
            } else if (c == ')' || c == ']' || c == '}') {
                char open = (c == ')') ? '(' : (c == ']') ? '[' : '{';
                if (top == 0 || stack[top - 1] != open) ok = false;
                else top--;
            }
        }
        if (top != 0) ok = false;

        System.out.println(ok ? "true" : "false");
    }
}
