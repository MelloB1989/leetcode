/*abstract
Scenic Flight Route
Scenic Flight Route

Problem Description

Ursula has won a contest, and the prize is a free flight trip that can consist of one or more flights through cities. Of course, Ursula wants to choose a trip that has as many cities as possible.

Ursula wants to fly from Uptown (City 1) to Downtown (City n) to visit the maximum number of cities. You are given the list of possible flights, and you know that there are no directed cycles in the flight network.

Input

The first line contains two integers n and m, separated by a space (2 ≤ n ≤ 10^5, 1 ≤ m ≤ 2 × 10^5): the number of cities and flights. The cities are numbered 1, 2, ..., n. City 1 is Uptown, and city n is Downtown.
Following that, there are m lines describing the flights. Each line has two integers a and b (1 ≤ a, b ≤ n): there is a flight from city a to city b. Each flight is a one-way flight.

Output

Print the maximum number of cities on the route. After this, print the cities in the order they will be visited. You can print any valid solution.

If there are no solutions, print "IMPOSSIBLE".

Constraints
2 ≤ n ≤ 10^5
1 ≤ m ≤ 2 × 10^5

Example
Input

5 5
1 2
2 5
1 3
3 4
4 5

Output

4
1 3 4 5

Note

In this example, the output represents the maximum number of cities on the route, and the order of cities visited. If there are no solutions, "IMPOSSIBLE" is printed.
*/
import java.util.*;

public class ScenicFlightRoute {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int m = sc.nextInt();

        // Build adjacency list
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            graph.add(new ArrayList<>());
        }

        int[] indegree = new int[n + 1];

        for (int i = 0; i < m; i++) {
            int a = sc.nextInt();
            int b = sc.nextInt();
            graph.get(a).add(b);
            indegree[b]++;
        }

        // Topological sort using Kahn's algorithm
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 1; i <= n; i++) {
            if (indegree[i] == 0) {
                queue.offer(i);
            }
        }

        List<Integer> topoOrder = new ArrayList<>();
        while (!queue.isEmpty()) {
            int curr = queue.poll();
            topoOrder.add(curr);

            for (int neighbor : graph.get(curr)) {
                indegree[neighbor]--;
                if (indegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }

        // DP to find longest path
        int[] dp = new int[n + 1];
        int[] parent = new int[n + 1];
        Arrays.fill(dp, -1);
        Arrays.fill(parent, -1);

        dp[1] = 1; // Starting city has 1 city (itself)

        // Process cities in topological order
        for (int city : topoOrder) {
            if (dp[city] == -1) continue; // Not reachable from city 1

            for (int neighbor : graph.get(city)) {
                if (dp[neighbor] < dp[city] + 1) {
                    dp[neighbor] = dp[city] + 1;
                    parent[neighbor] = city;
                }
            }
        }

        // Check if city n is reachable
        if (dp[n] == -1) {
            System.out.println("IMPOSSIBLE");
            return;
        }

        // Reconstruct path
        List<Integer> path = new ArrayList<>();
        int curr = n;
        while (curr != -1) {
            path.add(curr);
            curr = parent[curr];
        }
        Collections.reverse(path);

        // Output
        System.out.println(dp[n]);
        for (int i = 0; i < path.size(); i++) {
            if (i > 0) System.out.print(" ");
            System.out.print(path.get(i));
        }
        System.out.println();

        sc.close();
    }
}
