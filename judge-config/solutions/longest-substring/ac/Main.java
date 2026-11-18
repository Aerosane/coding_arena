// longest-substring - Accepted solution (Java). Expected verdict: AC.

import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String s = br.readLine();
        if (s == null) s = "";  // empty line / EOF -> empty string -> 0

        // Sliding window; last[c] = most recent index of char c (spaces count too).
        int[] last = new int[65536];
        Arrays.fill(last, -1);
        int best = 0, left = 0;
        for (int right = 0; right < s.length(); right++) {
            char c = s.charAt(right);
            if (last[c] >= left) left = last[c] + 1;
            last[c] = right;
            best = Math.max(best, right - left + 1);
        }

        System.out.println(best);
    }
}
