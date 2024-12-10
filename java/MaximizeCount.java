import java.io.*;
import java.util.*;

public class MaximizeCount {

    public static int count(int[] arr, int n) {
        int c = 0;
        int l = 0;
        int r = n - 1;
        int pStart = n;

        while (l <= r) {
            int mid = (l + r) / 2;
            if (arr[mid] > 0) {
                pStart = mid;
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }

        int neg = 0;
        int pos = 0;

        // Count negative numbers
        for (int i = 0; i < pStart; i++) {
            if (arr[i] < 0) {
                neg++;
            }
        }

        // Count positive numbers
        for (int i = pStart; i < n; i++) {
            if (arr[i] > 0) {
                pos++;
            }
        }

        c = Math.max(neg, pos);
        return c;
    }

    public static void main(String[] args) {
        int[] arr = { -3, -2, -1, -1, 0, 0, 1, 2 };
        System.out.println(count(arr, arr.length));
    }
}
