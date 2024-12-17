import java.util.Scanner;

public class MaximizeVowels {

    private static int countVowels(String s) {
        int count = 0;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                count++;
            }
        }
        return count;
    }

    public static int max(String s, int k) {
        int maxVowels = 0;
        for (int i = 0; i <= s.length() - k; i++) {
            String substring = s.substring(i, i + k);
            int vowelCount = countVowels(substring);
            maxVowels = Math.max(maxVowels, vowelCount);
        }

        return maxVowels;
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        String n = s.nextLine();
        int k = s.nextInt();
        s.close();

        System.out.println(max(n, k));
    }
}
