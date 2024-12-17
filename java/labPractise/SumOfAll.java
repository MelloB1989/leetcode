import java.math.BigInteger;
import java.util.Scanner;

public class SumOfAll {

    public static BigInteger sum(String s) {
        BigInteger sum = BigInteger.ZERO;
        int i = 0;

        while (i < s.length()) {
            if (s.charAt(i) < '0' || s.charAt(i) > '9') {
                i++;
                continue;
            }
            StringBuilder numStr = new StringBuilder();
            while (i < s.length() && s.charAt(i) >= '0' && s.charAt(i) <= '9') {
                numStr.append(s.charAt(i));
                i++;
            }
            sum = sum.add(new BigInteger(numStr.toString()));
        }

        return sum;
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        String n = s.nextLine();

        s.close();

        System.out.println(sum(n));
    }
}
