// merge-sorted-arrays - slow solution (Java): concatenate then selection sort, O(K^2).
// Correct on every official test (they are tiny); exceeds the time limit on large inputs (see run_solutions.py stress phase).

import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();
        for (String line; (line = br.readLine()) != null; ) sb.append(line).append(' ');
        StringTokenizer st = new StringTokenizer(sb.toString());

        int m = Integer.parseInt(st.nextToken());
        int n = Integer.parseInt(st.nextToken());
        int k = m + n;
        long[] v = new long[k];
        for (int i = 0; i < k; i++) v[i] = Long.parseLong(st.nextToken());

        for (int i = 0; i < k; i++) {  // selection sort
            int mn = i;
            for (int j = i + 1; j < k; j++)
                if (v[j] < v[mn]) mn = j;
            long t = v[i]; v[i] = v[mn]; v[mn] = t;
        }

        StringBuilder res = new StringBuilder();
        for (int i = 0; i < k; i++) {
            if (i > 0) res.append(' ');
            res.append(v[i]);
        }
        System.out.println(res);
    }
}
