import java.io.*;
import java.util.HashSet;
import java.util.Scanner;

public class Program {

    public static int sumOfLengths(int[] arr) {
        int n = arr.length;
        int sumOfLengths = 0;
        int start = 0;
        HashSet<Integer> set = new HashSet<>();

        for (int end = 0; end < n; end++) {
            // Remove elements till there are no duplicates
            while (set.contains(arr[end])) {
                set.remove(arr[start]);
                start++;
            }
            set.add(arr[end]);

            int windowLength = end - start + 1;
            sumOfLengths += (windowLength * (windowLength + 1)) / 2;
        }
        return sumOfLengths;
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
        int[] arr = new int[n];
        for (int i = 0; i < n; i++) {
            arr[i] = s.nextInt();
        }
        System.out.println(sumOfLengths(arr));
    }
}
