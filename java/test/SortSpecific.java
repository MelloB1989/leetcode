import java.io.*;
import java.util.*;

public class SortSpecific {

    public static void sort(int[] arr) {
        List<Integer> odds = new ArrayList<>();
        List<Integer> evens = new ArrayList<>();

        int n = arr.length;

        for (int num : arr) {
            if (num % 2 != 0) {
                odds.add(num);
            } else {
                evens.add(num);
            }
        }

        Collections.sort(odds, Collections.reverseOrder());

        Collections.sort(evens);

        List<Integer> result = new ArrayList<>();
        result.addAll(odds);
        result.addAll(evens);

        for (int num : result) {
            System.out.print(num + " ");
        }
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        int n = s.nextInt();
        int[] arr = new int[n];

        for (int i = 0; i < n; i++) {
            arr[i] = s.nextInt();
        }

        sort(arr);
    }
}
