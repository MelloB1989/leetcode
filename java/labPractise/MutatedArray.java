import java.io.*;
import java.util.*;

public class MutatedArray {

    public static int mutate(int[] arr, int target) {
        int max = 0;
        for (int num : arr) {
            max = Math.max(max, num);
        }

        int left = 0, right = max;
        int minDifference = Integer.MAX_VALUE;
        int bestValue = -1;

        while (left <= right) {
            int mid = (left + right) / 2;
            int currentSum = calculateSum(arr, mid);
            int diff = Math.abs(currentSum - target);

            if (
                diff < minDifference ||
                (diff == minDifference && mid < bestValue)
            ) {
                minDifference = diff;
                bestValue = mid;
            }

            if (currentSum < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return bestValue;
    }

    private static int calculateSum(int[] arr, int value) {
        int sum = 0;
        for (int num : arr) {
            sum += Math.min(num, value);
        }
        return sum;
    }

    public static void main(String[] args) {
        // Scanner s = new Scanner(System.in);

        // int n = s.nextInt();
        // int[] arr = new int[n];

        // for (int i = 0; i < n; i++) {
        //     arr[i] = s.nextInt();
        // }
        int[] arr = { 4, 9, 3 };
        int k = 10;
        System.out.println(mutate(arr, k));
    }
}
