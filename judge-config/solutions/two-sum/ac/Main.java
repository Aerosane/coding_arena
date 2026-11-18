// two-sum - Accepted solution (Java). Expected verdict: AC.

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

        // One pass with a hash map: value -> first index it appeared at.
        HashMap<Long, Integer> seen = new HashMap<>();
        for (int j = 0; j < n; j++) {
            Integer i = seen.get(target - a[j]);
            if (i != null) {
                System.out.println(i + " " + j);
                return;
            }
            seen.putIfAbsent(a[j], j);
        }
    }
}
