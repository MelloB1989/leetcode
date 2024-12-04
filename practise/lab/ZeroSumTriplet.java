import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class ZeroSumTriplet {

    public static void findTriplet(int[] nums, int size) {
        for (int i = 0; i < size; i++) {
            int left = nums[i + 1];
            int right = nums[size - 1];
            int curr = nums[i];
            int sum = curr + left + right;
            if (sum == 0) {}
        }
    }

    public static void main(String[] args) {
        int[] arr = { 0, -1, 2, -3, 1 };
        int n = 5;
        findTriplet(arr, n);
    }
    // public static void main(String[] args) {
    //     Scanner s = new Scanner(System.in);
    //     int n = s.nextInt();
    //     int[] arr = new int[n];
    //     for (int i = 0; i < n; i++) {
    //         arr[i] = s.nextInt();
    //     }
    //     sort(arr, n);
    // }
}
