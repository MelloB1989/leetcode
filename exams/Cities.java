/*
Minimizing Library Access Costs

The task is to determine the minimum cost required to provide library access to all citizens of HackerLand. In HackerLand, there are n cities numbered from 1 to n. Initially, there are no libraries, and the cities are not connected. However, bidirectional roads can be built between any city pair listed in cities. A citizen has access to a library if:

Their city contains a library.
They can travel by road from their city to a city containing a library.
For each query:

The input specifies a map of HackerLand along with the cost to build libraries (c_lib) and repair roads (c_road).
The goal is to find the minimum cost to make libraries accessible to all citizens by building libraries and connecting cities through roads.

int n: the number of cities
int c_lib: the cost to build a library
int c_road: the cost to repair a road
int cities[m][2]: each cities[i] contains two integers representing cities that can be connected by a new road
Returns:

int: the minimal cost to provide library access to all citizens

Input Format:

The first line contains an integer q denoting the number of queries.

For each query:

The first line contains four space-separated integers: n, m, c_lib, and c_road.
Each of the next m lines contains two space-separated integers, u[i] and v[i], representing bidirectional roads between cities u[i] and v[i].

Constraints:

1 <= q <= 10
1 <= n <= 10^5
0 <= m <= min(10^5, n*(n-1)/2)
1 <= c_road, c_lib <= 10^5
1 <= u[i], v[i] <= n
Each road connects two distinct cities

Sample Input:

2
3 3 2 1
1 2
3 1
2 3
6 6 2 5
1 3
3 4
2 4
1 2
2 3
5 6

output
4
12

Explanation (for first query)

HackerLand contains n=3  cities and can be connected by m=3 bidirectional roads. The price of building a library is 2 and the price for repairing a road is 1 .

The cheapest way to make libraries accessible to all is to:

Build a library in city 1 at a cost of 2.
Build the road between cities 1 and 2 at a cost of 1.
Build the road between cities 2 and 3 at a cost of 1.
This gives a total cost of 4. Note that the road between cities 3 and 1 does not need to be built because each is connected to city 2.
*/
import java.util.*;

public class Main {

    public static long minimumCost(int n, int cLib, int cRoad, int[][] cities) {
        if (cLib <= cRoad) {
            return (long) n * cLib;
        }

        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            graph.add(new ArrayList<>());
        }

        for (int[] road : cities) {
            graph.get(road[0]).add(road[1]);
            graph.get(road[1]).add(road[0]);
        }

        boolean[] visited = new boolean[n + 1];
        long totalCost = 0;

        for (int i = 1; i <= n; i++) {
            if (!visited[i]) {
                int componentSize = dfs(graph, visited, i);
                totalCost += cLib + (long) (componentSize - 1) * cRoad;
            }
        }

        return totalCost;
    }

    private static int dfs(
        List<List<Integer>> graph,
        boolean[] visited,
        int city
    ) {
        visited[city] = true;
        int size = 1;

        for (int neighbor : graph.get(city)) {
            if (!visited[neighbor]) {
                size += dfs(graph, visited, neighbor);
            }
        }

        return size;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int q = scanner.nextInt();

        for (int query = 0; query < q; query++) {
            int n = scanner.nextInt();
            int m = scanner.nextInt();
            int cLib = scanner.nextInt();
            int cRoad = scanner.nextInt();

            int[][] cities = new int[m][2];
            for (int i = 0; i < m; i++) {
                cities[i][0] = scanner.nextInt();
                cities[i][1] = scanner.nextInt();
            }

            System.out.println(minimumCost(n, cLib, cRoad, cities));
        }

        scanner.close();
    }
}
