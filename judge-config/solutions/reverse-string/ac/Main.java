// reverse-string - Accepted solution (Java). Expected verdict: AC.

import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        // Test data is one raw line (spaces are significant), e.g. "hello world".
        String s = br.readLine();
        if (s == null) s = "";

        String result = new StringBuilder(s).reverse().toString();

        System.out.println(result);
    }
}
