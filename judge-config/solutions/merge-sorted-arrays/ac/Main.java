// merge-sorted-arrays - Accepted solution (Java). Expected verdict: AC.

import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();
        for (String line; (line = br.readLine()) != null; ) sb.append(line).append(' ');
        // Token-based parsing: an array line may be empty when m or n is 0.
        StringTokenizer st = new StringTokenizer(sb.toString());

        int m = Integer.parseInt(st.nextToken());
        int n = Integer.parseInt(st.nextToken());
        long[] a = new long[m];
        long[] b = new long[n];
        for (int i = 0; i < m; i++) a[i] = Long.parseLong(st.nextToken());
        for (int j = 0; j < n; j++) b[j] = Long.parseLong(st.nextToken());

        long[] out = new long[m + n];
        int i = 0, j = 0, k = 0;
        while (i < m && j < n) out[k++] = (a[i] <= b[j]) ? a[i++] : b[j++];
        while (i < m) out[k++] = a[i++];
        while (j < n) out[k++] = b[j++];

        StringBuilder res = new StringBuilder();
        for (int x = 0; x < out.length; x++) {
            if (x > 0) res.append(' ');
            res.append(out[x]);
        }
        System.out.println(res);
    }
}
