import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class FreqSort {

    public static void sort(int[] nums, int size) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < size; i++) {
            map.put(nums[i], map.getOrDefault(nums[i], 0) + 1);
        }

        List<Integer> unique = new ArrayList<>(map.keySet());

        Collections.sort(unique, (a, b) -> {
            int freqCompare = map.get(b).compareTo(map.get(a));
            if (freqCompare == 0) {
                return a.compareTo(b);
            }
            return freqCompare;
        });

        List<Integer> result = new ArrayList<>();
        for (int num : unique) {
            int freq = map.get(num);
            for (int j = 0; j < freq; j++) {
                result.add(num);
            }
        }
        for (int i = 0; i < result.size(); i++) {
            System.out.print(" " + result.get(i));
        }
    }

    public static void main(String[] args) {
        int[] arr = { 2, 5, 2, 8, 5, 6, 8, 8 };
        int n = 8;
        sort(arr, n);
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
