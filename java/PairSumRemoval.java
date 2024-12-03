import java.io.*;
import java.util.HashMap;
import java.util.Scanner;

public class PairSumRemoval {

    public static int maxOperations(int[] nums, int k) {
        HashMap<Integer, Integer> map = new HashMap<>();
        int count = 0;

        for (int num : nums) {
            int complement = k - num;
            if (map.getOrDefault(complement, 0) > 0) {
                count++;
                map.put(complement, map.get(complement) - 1);
            } else {
                map.put(num, map.getOrDefault(num, 0) + 1);
            }
        }

        return count;
    }

    public static void main(String[] args) {
        int[] arr = { 3, 1, 3, 4, 3 };
        int k = 6;
        System.out.println(maxOperations(arr, k));
    }
}
