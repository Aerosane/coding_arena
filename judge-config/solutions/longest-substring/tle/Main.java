// longest-substring - slow solution (Java): scan every start to the end with no early break, O(N^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String s = br.readLine();
        if (s == null) s = "";

        int n = s.length();
        int[] stamp = new int[65536];
        int best = 0;
        for (int i = 0; i < n; i++) {
            // Scan from i to the end every time, even after a repeat was seen.
            boolean dup = false;
            int len = 0;
            for (int j = i; j < n; j++) {
                char c = s.charAt(j);
                if (!dup) {
                    if (stamp[c] == i + 1) dup = true;
                    else { stamp[c] = i + 1; len++; }
                }
            }
            best = Math.max(best, len);
        }

        System.out.println(best);
    }
}
