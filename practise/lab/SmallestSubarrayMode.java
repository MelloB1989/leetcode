import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class SmallestSubarrayMode {

    public static void findMode(int[] nums, int size) {
        HashMap<Integer, Integer> freqMap = new HashMap<>();
        HashMap<Integer, List<Integer>> occ = new HashMap<>();

        for (int i = 0; i < size; i++) {
            List<Integer> oc = occ.getOrDefault(nums[i], new ArrayList<>());
            oc.add(i);
            occ.put(nums[i], oc);
            freqMap.put(nums[i], freqMap.getOrDefault(nums[i], 0) + 1);
        }

        int max = 0;
        for (int freq : freqMap.values()) {
            if (freq > max) {
                max = freq;
            }
        }

        int minLen = size;

        for (int o : new ArrayList<>(freqMap.keySet())) {
            int freq = freqMap.get(o);
            if (freq == max) {
                List<Integer> l = occ.get(o);
                int len = l.get(l.size() - 1) - l.get(0) + 1;
                if (len < minLen) {
                    minLen = len;
                }
            }
        }
        System.out.print(minLen);
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
        int[] arr = new int[n];
        for (int i = 0; i < n; i++) {
            arr[i] = s.nextInt();
        }
        findMode(arr, n);
    }
}
