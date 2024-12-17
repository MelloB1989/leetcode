import java.math.BigInteger;
import java.util.Scanner;

public class ApplesInGarden {

    public static int findMinPerimeter(long n) {
        long low = 0;
        long high = 2_000_000_000;

        while (low < high) {
            long mid = low + (high - low) / 2;
            BigInteger apples = f(mid);

            if (apples.compareTo(BigInteger.valueOf(n)) >= 0) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }

        long D = high;
        long perimeter = D * 8;
        return (int) perimeter;
    }

    private static BigInteger f(long D) {
        BigInteger BD = BigInteger.valueOf(D);
        BigInteger D3 = BD.multiply(BD).multiply(BD);
        BigInteger D2 = BD.multiply(BD);

        BigInteger result = D3.shiftLeft(2)
            .add(D2.multiply(BigInteger.valueOf(6)))
            .add(BD.shiftLeft(1));
        return result;
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        long n = s.nextLong();

        s.close();

        System.out.println(findMinPerimeter(n));
    }
}
