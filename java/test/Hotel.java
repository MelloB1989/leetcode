/**
Hotel Accommodation Challenge

In a hotel with \(M) floors and various room capacities per floor, \(N) people arrive for accommodation. Each floor has identical rooms that can accommodate specific numbers of people.

Your task is to count the different ways to accommodate
N
 people in the hotel rooms according to the provided rules.

Input Format
- First line: Two integers \(M) and \(N) denoting the number of floors and the number of people, respectively (\(1<=N*M<=1e6))
- Second line: \(M) space-separated integers denoting the capacity of each floor, where the \(i)th integer represents the capacity of the \(i)th floor (\(1<=C[i]<=1e6))

Output Format
Print the number of different ways of accommodating people. As the count could be large, output the answer modulo
10
9
+
7
.

Constraints
- All \(C[i]) values are different.
- All \(M) values are different.

Sample Input
3 5
1 2 3

Sample Output
5

Explanation


You can assign the people as follows:
- (1,1,1,1,1): Assign each of the 5 people to rooms on the first floor.
- (1,1,1,2): Assign 3 people to rooms on the first floor and 2 people to rooms on the second floor.
- (1,1,3): Assign 2 people to rooms on the first floor and 3 people to a room on the third floor.
- (1,2,2): Assign 4 people to rooms on the second floor, with each room accommodating 2 people, and 1 person to a room on the first floor.
- (2,3): Assign 2 people to a room on the second floor and 3 people to a room on the third floor.
*/
import java.util.Scanner;

public class Hotel {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int M = scanner.nextInt();
        int N = scanner.nextInt();

        int[] capacities = new int[M];
        for (int i = 0; i < M; i++) {
            capacities[i] = scanner.nextInt();
        }

        int MOD = 1000000007;

        int[] dp = new int[N + 1];
        dp[0] = 1;

        for (int capacity : capacities) {
            for (int j = capacity; j <= N; j++) {
                dp[j] = (dp[j] + dp[j - capacity]) % MOD;
            }
        }

        System.out.println(dp[N]);

        scanner.close();
    }
}
