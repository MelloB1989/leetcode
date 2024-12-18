import java.io.*;
import java.util.*;

public class SentencePalindrome {

    public static String check(String s) {
        String ans = "Yes";
        String n = "";

        for (int i = 0; i < s.length(); i++) {
            if (
                (s.charAt(i) >= 65 && s.charAt(i) <= 90) ||
                (s.charAt(i) >= 97 && s.charAt(i) <= 122)
            ) {
                n += s.charAt(i);
            }
        }

        int left = 0;
        int right = n.length() - 1;
        while (left <= right) {
            if (n.charAt(left) != n.charAt(right)) {
                ans = "No";
                break;
            }
            left++;
            right--;
        }

        return ans;
    }

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        String s = input.nextLine();
        input.close();
        System.out.println(check(s));
    }
}
