/*
Building Cross-Team Friendships
Building Cross-Team Friendships

You've established a school with N students, each of whom is designated to either Team Red or Team Blue. However, you aspire for students to foster friendships across teams. Each student has a set number M, representing the count of friends they should have from the opposite team.

Your task is to ascertain if it's feasible for each student to achieve the specified number of friends from the opposing team.

Input Format:

The first line contains a number T denoting the number of test cases.

For each test case:

The first line contains a positive integer N, denoting the number of students.
The second line contains an integer M, representing the number of friends that each student should have from the opposite team.

Output Format:

For each test case, print "Yes" or "No" in a new line based on the possibility of a student achieving the specified number of friends represented by M.

Constraints:

1 <= T <= 2*10^5
1 <= N <= 10^5
0 <= M <= 10^5

Sample Input:

1

2 1

Sample Output:


Yes

Explanation:

For instance, with n=2 and m=1, there are 2 students in the school, where student 1 belongs to Team Red and student 2 belongs to Team Blue. If they become friends, their requirements will be fulfilled as they both require a friend from the opposite team. Thus, the output is "Yes".

Note: The task is to determine whether it's possible for each student to have the specified number of friends from the opposing team based on the given constraints.
*/
import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        StringBuilder result = new StringBuilder();

        while (T-- > 0) {
            int N = sc.nextInt();
            int M = sc.nextInt();

            int R = N / 2;
            int B = N - R;

            if (M <= R && M <= B) {
                result.append("Yes\n");
            } else {
                result.append("No\n");
            }
        }

        System.out.print(result);
    }
}
