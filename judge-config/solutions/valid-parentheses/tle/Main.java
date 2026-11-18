// valid-parentheses - slow solution (Java): repeated pair removal, O(N^2) on deeply nested input.
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String line = br.readLine();
        StringBuilder s = new StringBuilder(line == null ? "" : line.trim());

        // Keep deleting one adjacent matching pair per scan until nothing changes.
        String[] pairs = {"()", "[]", "{}"};
        boolean changed = true;
        while (changed) {
            changed = false;
            for (String p : pairs) {
                int pos = s.indexOf(p);
                if (pos >= 0) {
                    s.delete(pos, pos + 2);
                    changed = true;
                }
            }
        }

        System.out.println(s.length() == 0 ? "true" : "false");
    }
}
