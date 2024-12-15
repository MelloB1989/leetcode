import java.io.*;
import java.util.*;

public class ApplesInAGarden {

    public static int findApples(int k) {
        return 4 * k * k + 6 * k * k + 2 * k;
    }

    public static int findMinPerimeter(long n) {
        int right = 1;
        int left = 0;

        while (findApples(right) < n) {
            right *= 2;
        }

        while (left < right) {
            int mid = (left + (right - left)) / 2;
            if (findApples(mid) < n) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        return left * 8;
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        long n = s.nextLong();

        System.out.println(findMinPerimeter(n));
    }
}
