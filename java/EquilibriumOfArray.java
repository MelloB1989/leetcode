import java.io.*;
import java.util.Scanner;

public class EquilibriumOfArray {

    public static int eq(int[] arr) {
        int n = arr.length;
        for (int i = 0; i < n; i++) {
            int leftSum = 0;
            int rightSum = 0;
            for (int j = 0; j < i; j++) {
                leftSum += arr[j];
            }
            for (int j = i + 1; j < n; j++) {
                rightSum += arr[j];
            }
            System.out.println(leftSum + " " + rightSum);
            if (leftSum == rightSum) {
                return i + 1;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        int[] arr = { 1, 2, 3, 4, 8, 10 };
        System.out.println(eq(arr));
    }
}
