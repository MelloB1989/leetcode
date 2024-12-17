import java.io.*;
import java.util.*;

public class LowerBound {

    public static int getLowerBound(int[] arr, int k) {
        int n = arr.length;

        Arrays.sort(arr);

        int left = 0;
        int right = n - 1;
        if (k > arr[n - 1]) {
            return -1;
        }
        while (left < right) {
            int mid = (left + right) / 2;
            if (arr[mid] == k) {
                return k;
            } else if (arr[mid] > k) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return arr[left];
    }

    public static void main(String[] args) {
        // Scanner s = new Scanner(System.in);

        // int n = s.nextInt();
        // int[] arr = new int[n];

        // for (int i = 0; i < n; i++) {
        //     arr[i] = s.nextInt();
        // }
        int[] arr = { 6, 7, 3, 4, 5, 1, 2 };
        int k = 3;
        System.out.println(getLowerBound(arr, k));
    }
}
