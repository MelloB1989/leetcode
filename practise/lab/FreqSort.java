import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class FreqSort {

    public static void sort(int[] nums, int size) {
        HashMap<Integer, Integer> map = new HashMap<>();
        List<Integer> unique = new ArrayList<>();
        List<Integer> result = new ArrayList<>();

        for (int i = 0; i < size; i++) {
            map.put(nums[i], map.getOrDefault(nums[i], 0) + 1);
            if (!unique.contains(nums[i])) {
                unique.add(nums[i]);
            }
        }

        HashMap<Integer, List<Integer>> frequencies = new HashMap<>();
        for (int key : unique) {
            int freq = map.get(key);
            if (!frequencies.containsKey(freq)) {
                frequencies.put(freq, new ArrayList<>());
            }
            frequencies.get(freq).add(key);
        }

        List<Integer> sortedFrequencies = new ArrayList<>(frequencies.keySet());
        Collections.sort(sortedFrequencies);
        sortedFrequencies = sortedFrequencies.reversed();

        for (int freq : sortedFrequencies) {
            List<Integer> elements = frequencies.get(freq);
            for (int element : elements) {
                for (int i = 0; i < freq; i++) {
                    result.add(element);
                }
            }
        }

        for (int i = 0; i < result.size(); i++) {
            System.out.print(result.get(i) + " ");
        }
    }

    public static void main(String[] args) {
        int[] arr = { 2, 5, 2, 6, -1, 9999999, 5, 8, 8, 8 };
        int n = 10;
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
