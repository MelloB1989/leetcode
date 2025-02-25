/**
You're given an (n times n) grid with some squares marked as traps. Moving right or down is allowed, but you cannot step onto a square with a trap. Your goal is to compute the number of possible paths from the upper-left square to the lower-right square.

Input:

The first line contains an integer (n), denoting the grid's size.

Following that are (n) lines, each containing (n) characters. A period (.) represents an empty cell, while an asterisk (*) denotes a trap.

Output:

Print the count of paths modulo (10^9 + 7).

Constraints:

1<=N<=1000

Example:


Input:


4
....
.*..
...*
*...

Output:

3
*/
import java.util.Scanner;

public class Traps {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        scanner.nextLine(); // Consume the newline character

        char[][] grid = new char[n][n];
        for (int i = 0; i < n; i++) {
            grid[i] = scanner.nextLine().toCharArray();
        }

        System.out.println(countPaths(grid, n));
        scanner.close();
    }

    private static int countPaths(char[][] grid, int n) {
        int MOD = 1000000007;
        int[][] dp = new int[n][n];

        // Initialize the starting point
        if (grid[0][0] == '.') {
            dp[0][0] = 1;
        }

        // Fill the dp array
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '*') {
                    dp[i][j] = 0; // No paths through a trap
                } else {
                    if (i > 0) {
                        dp[i][j] = (dp[i][j] + dp[i - 1][j]) % MOD;
                    }
                    if (j > 0) {
                        dp[i][j] = (dp[i][j] + dp[i][j - 1]) % MOD;
                    }
                }
            }
        }

        return dp[n - 1][n - 1];
    }
}
