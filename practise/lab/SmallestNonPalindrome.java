import java.io.*;
import java.util.*;

public class SmallestNonPalindrome {

    public static String smallest(String s) {
        String result = "";
        char replacement = 'a';
        // boolean replaced = false;

        for (int i = 0; i < s.length(); i++) {
            if (replacement < s.charAt(i)) {
                StringBuilder sb = new StringBuilder(s);
                sb.setCharAt(i, replacement);
                result = sb.toString();
                break;
            }
        }

        return result;
    }

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        String s = input.nextLine();
        input.close();
        System.out.println(smallest(s));
    }
}
