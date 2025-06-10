/*
Topological Sort
Problem

You are given a Directed Acyclic Graph (DAG) with N vertices and M edges. Your task is to print the topological sorting of the vertices.

Input:

The first line contains two space-separated integers, N and M, denoting the number of vertices and edges in the graph.
Each of the following M lines contains two space-separated integers, X and Y, indicating a directed edge from vertex X to vertex Y.

Output:

Print N space-separated integers representing the topological sort of the vertices. If there are multiple possible orderings, print the lexicographically smallest one.

Constraints:

1 <= N <= 10
1 <= M <= 20
1 <= X, Y <= N

Sample Input:

5 6
1 2
1 3
2 3
2 4
3 4
3 5

Sample Output: 1 2 3 4 5

Explanation:
In the given sample, the topological sorting of the vertices is {1, 2, 3, 4, 5}. This represents a valid order such that for every directed edge (X, Y), vertex X comes before vertex Y in the ordering. If there are multiple valid orderings, print the lexicographically smallest one.
*/
import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int m = sc.nextInt();
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i <= n; i++) graph.add(new ArrayList<>());
        int[] inDegree = new int[n + 1];
        for (int i = 0; i < m; i++) {
            int x = sc.nextInt();
            int y = sc.nextInt();
            graph.get(x).add(y);
            inDegree[y]++;
        }
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        for (int i = 1; i <= n; i++) {
            if (inDegree[i] == 0) pq.add(i);
        }
        List<Integer> topoOrder = new ArrayList<>();
        while (!pq.isEmpty()) {
            int curr = pq.poll();
            topoOrder.add(curr);
            for (int neighbor : graph.get(curr)) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) pq.add(neighbor);
            }
        }
        for (int i = 0; i < topoOrder.size(); i++) {
            System.out.print(topoOrder.get(i));
            if (i < topoOrder.size() - 1) System.out.print(" ");
        }
    }
}
