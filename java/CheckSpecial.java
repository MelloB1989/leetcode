//Please accept the inputs from the console and print the output
import java.io.*;
import java.util.Scanner;

//********Please do not remove/change Class Name *********
public class CheckSpecial {

    private static boolean check(String o, String g) {
        if (o.length() < 2 || o.length() != g.length()) {
            return false;
        }
        int a = -1;
        int b = -1;
        for (int i = 0; i < o.length(); i++) {
            char originalChar = o.charAt(i);
            char goalChar = g.charAt(i);
            if (originalChar != goalChar) {
                if (a == -1) {
                    a = i;
                } else if (b == -1) {
                    b = i;
                } else {
                    return false;
                }
            }
        }
        if (a != -1 && b != -1) {
            return o.charAt(a) == g.charAt(b) && o.charAt(b) == g.charAt(a);
        }
        return false;
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String original = in.nextLine();
        String goal = in.nextLine();
        boolean res = check(original, goal);
        if (res) {
            System.out.println("1");
        } else {
            System.out.println("0");
        }
    }
}
