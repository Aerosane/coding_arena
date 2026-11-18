// reverse-string - slow solution (Java): repeated front insertion, O(N^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String s = br.readLine();
        if (s == null) s = "";

        StringBuilder r = new StringBuilder();
        for (int i = 0; i < s.length(); i++) r.insert(0, s.charAt(i));  // shifts everything built so far

        System.out.println(r);
    }
}
