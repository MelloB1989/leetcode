/*abstractShortest Paths from Starting Node in an Undirected Graph

Problem Statement:

You are given an undirected graph and a starting node. Determine the lengths of the shortest paths from the starting node to all other nodes in the graph. If a node is unreachable, its distance is -1. Nodes are numbered consecutively from 1 to n, and edges have varying distances or lengths.

For instance, consider a graph with 5 nodes:

mathematica
Copy code
Begin End Weight
1     2   5
2     3   6
3     4   2
1     3   15
Starting at node 1, the shortest path to node 2 is direct with a distance of 5. Moving from node 1 to node 3, there are two possible paths: 1->2->3 with a distance of 5 + 6 = 11 or 1->3 with a distance of 15. The shorter path is 1->2->3, totaling 11. Continuing from node 1 to node 4, the shortest path goes through node 3: 1->2->3->4 with a distance of 11 + 2 = 13. Node 5 is unreachable, so the distance to it is -1.

The distances to all nodes, in increasing order of node number and excluding the starting node, are 5, 11, 13, -1.

Input Format:
The first line contains an integer t, the number of test cases.

For each test case:

The first line contains two space-separated integers n and m, representing the number of nodes and edges in the graph.
Each of the next m lines contains three space-separated integers x, y, and r, denoting the starting and ending nodes of an edge and the length of the edge.
The last line of each test case contains an integer s, indicating the starting position.

Constraints:

1 <= t <= 10
2 <= n <= 3000
1 <= m <= N*(N-1)/2
1 <= x, y, s <= N
1 <= r <= 1e5

Multiple edges between the same pair of nodes with different weights are to be considered independently.

Output Format:

For each of the t test cases, print a single line consisting of n-1 space-separated integers. These integers denote the shortest distances from the starting position s to the n-1 nodes in increasing order of their labels (excluding s). For unreachable nodes, print -1.

Sample Input:

1
4 4
1 2 24
1 4 20
3 1 3
4 3 12
1

Sample Output:

24 3 15
*/
import java.util.*;

public class Main {

    static class Edge {

        int to, weight;

        Edge(int to, int weight) {
            this.to = to;
            this.weight = weight;
        }
    }

    static class Node implements Comparable<Node> {

        int vertex, distance;

        Node(int vertex, int distance) {
            this.vertex = vertex;
            this.distance = distance;
        }

        @Override
        public int compareTo(Node other) {
            return Integer.compare(this.distance, other.distance);
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int t = sc.nextInt();

        while (t-- > 0) {
            int n = sc.nextInt();
            int m = sc.nextInt();

            List<List<Edge>> graph = new ArrayList<>();
            for (int i = 0; i <= n; i++) {
                graph.add(new ArrayList<>());
            }

            for (int i = 0; i < m; i++) {
                int x = sc.nextInt();
                int y = sc.nextInt();
                int r = sc.nextInt();

                graph.get(x).add(new Edge(y, r));
                graph.get(y).add(new Edge(x, r));
            }

            int start = sc.nextInt();

            int[] distances = new int[n + 1];
            Arrays.fill(distances, Integer.MAX_VALUE);
            distances[start] = 0;

            PriorityQueue<Node> pq = new PriorityQueue<>();
            pq.offer(new Node(start, 0));

            while (!pq.isEmpty()) {
                Node current = pq.poll();
                int u = current.vertex;
                int currentDist = current.distance;

                if (currentDist > distances[u]) {
                    continue;
                }

                for (Edge edge : graph.get(u)) {
                    int v = edge.to;
                    int weight = edge.weight;
                    int newDist = distances[u] + weight;

                    if (newDist < distances[v]) {
                        distances[v] = newDist;
                        pq.offer(new Node(v, newDist));
                    }
                }
            }

            List<Integer> result = new ArrayList<>();
            for (int i = 1; i <= n; i++) {
                if (i != start) {
                    if (distances[i] == Integer.MAX_VALUE) {
                        result.add(-1);
                    } else {
                        result.add(distances[i]);
                    }
                }
            }

            for (int i = 0; i < result.size(); i++) {
                if (i > 0) System.out.print(" ");
                System.out.print(result.get(i));
            }
            System.out.println();
        }

        sc.close();
    }
}
