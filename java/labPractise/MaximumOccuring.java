import java.util.HashMap;
import java.util.Scanner;

public class MaximumOccuring {

    public static char max(String s) {
        HashMap<Character, Integer> map = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
        }
        int max = 0;
        char c = '\0';
        for (char occ : map.keySet()) {
            int count = map.get(occ);
            if (count > max || (count == max && occ < c)) {
                max = count;
                c = occ;
            }
        }
        return c;
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        String n = s.nextLine();

        s.close();

        System.out.println(max(n));
    }
}
