// two-sum - slow solution (Java): brute force over all pairs, O(N^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();
        for (String line; (line = br.readLine()) != null; ) sb.append(line).append(' ');
        StringTokenizer st = new StringTokenizer(sb.toString());

        int n = Integer.parseInt(st.nextToken());
        long[] a = new long[n];
        for (int i = 0; i < n; i++) a[i] = Long.parseLong(st.nextToken());
        long target = Long.parseLong(st.nextToken());

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (a[i] + a[j] == target) {
                    System.out.println(i + " " + j);
                    return;
                }
            }
        }
    }
}
