/**
A new AI has been designed to read horizontal strips of memory. The memory strip is divided into blocks numbered 1, 2, 3, and so on. Certain blocks contain valuable data marked as good memory points, and these points are given in an array 'memory'. The AI can only read at these specific points.

The AI's goal is to reach the last point of good memory, moving only forward and reading the memory sequentially. It can jump from one good memory point to another following specific rules:
- After moving from one point to another, it can jump to points L-1, L, or L+1, where the length of its last jump was L.
- The AI begins at point 0 (always a good memory point), and the first move must be to the first good memory point.

Determine whether the AI can reach the last memory point. The good memory points are given in ascending order.

Input Format:

- The first line contains an integer T, representing the number of test cases.
- For each test case:
  - The first line contains an integer n, denoting the number of good memory points.
  - The second line contains n space-separated integers in ascending order, representing the positions of the good memory points.

Output Format:

For each test case, print "YES" if the AI can read the last memory point, otherwise print "NO".

Constraints:
- 1 <= T <= 10
- 1 <= n <= 1000
- 1 <= memory[i] <= 10^9
- memory[i] < memory[i+1] (The memory points are in ascending order)

Sample Input:

2
6
1 2 3 4 5 6
2
1 11


Sample Output:

YES
NO


Explanation:
- In the first test case, the AI starts at 0 and successfully jumps to reach the last memory point at 6.
- In the second test case, despite reaching the first memory point at 1, the required jump distance of 10 to reach 11 is not possible. Hence, the AI cannot read the last memory point.
*/
import java.util.*;

public class AIRead {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int T = scanner.nextInt();

        for (int t = 0; t < T; t++) {
            int n = scanner.nextInt();
            int[] memory = new int[n];

            for (int i = 0; i < n; i++) {
                memory[i] = scanner.nextInt();
            }

            if (canReachLastMemoryPoint(memory)) {
                System.out.println("YES");
            } else {
                System.out.println("NO");
            }
        }

        scanner.close();
    }

    private static boolean canReachLastMemoryPoint(int[] memory) {
        int n = memory.length;
        if (n == 1) return true; // If there's only one point, we're already there

        // Use a set to keep track of reachable points
        Set<Integer> reachable = new HashSet<>();
        reachable.add(0); // Start at point 0

        for (int i = 0; i < n; i++) {
            int currentPoint = memory[i];
            Set<Integer> newReachable = new HashSet<>();

            for (int lastJump : reachable) {
                int jumpLength = currentPoint - lastJump;

                if (jumpLength > 0) {
                    newReachable.add(currentPoint);
                    newReachable.add(currentPoint + jumpLength - 1);
                    newReachable.add(currentPoint + jumpLength);
                    newReachable.add(currentPoint + jumpLength + 1);
                }
            }

            reachable = newReachable;

            if (reachable.contains(memory[n - 1])) {
                return true;
            }
        }

        return false;
    }
}
