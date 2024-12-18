import java.io.*;
import java.util.*;

public class LongestSubstring {

    public static int longest(String s, int k) {
        int maxLen = -1;

        HashMap<Character, Integer> charFreq = new HashMap<>();

        int count = 0;
        int start = 0;

        for (int end = 0; end < s.length(); end++) {
            char curr = s.charAt(end);
            charFreq.put(curr, charFreq.getOrDefault(curr, 0) + 1);
            if (charFreq.get(curr) == 1) {
                count++;
            }
            // System.out.println(curr + " " + count);
            while (count > k) {
                char startChar = s.charAt(start);
                charFreq.put(startChar, charFreq.get(startChar) - 1);
                if (charFreq.get(startChar) == 0) {
                    count -= 1;
                }
                start += 1;
            }
            if (count == k) {
                maxLen = Math.max(maxLen, end - start + 1);
            }
        }
        return maxLen;
    }

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        String s = input.nextLine();
        int k = input.nextInt();
        input.close();
        System.out.println(longest(s, k));
    }
}
