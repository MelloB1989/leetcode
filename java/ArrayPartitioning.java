import java.io.*;
import java.util.*;

public class ArrayPartitioning {

    public static int parts(int[] arr) {
        int diff = Integer.MAX_VALUE;
        int n = arr.length;
        int[] preSum = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            preSum[i] = preSum[i - 1] + arr[i - 1];
        }

        for (int i = 1; i < n - 2; i++) {
            for (int j = i + 1; j < n - 1; j++) {
                for (int k = j + 1; k < n; k++) {
                    int w = preSum[i];
                    int x = preSum[j] - preSum[i];
                    int y = preSum[k] - preSum[j];
                    int z = preSum[n] - preSum[k];

                    int maxSum = Math.max(Math.max(w, x), Math.max(y, z));
                    int minSum = Math.min(Math.min(w, x), Math.min(y, z));

                    // System.out.println(maxSum + " " + minSum);

                    diff = Math.min(diff, maxSum - minSum);
                }
            }
        }
        return diff;
    }

    public static void main(String[] args) {
        int[] arr = { 4, 2, 2, 5, 1 };
        System.out.println(parts(arr));
    }
}
