/*abstract
Collaborative Partnerships (Collabpartnership)
Collaborative Partnerships

In the realm of Crypta, renowned for its challenging contests, a unique collaboration system governs team interactions. Teams can work together under specific conditions: if Team X can collaborate with Team Y, and Team Y can collaborate with Team Z, then it implies that Team X can also collaborate with Team Z, and vice versa.

During the contest, Crypta makes two types of announcements:

J X Y: Indicates that Team X can collaborate with Team Y.
? X Y: Query to determine if Team X and Team Y can work together.
For each second type of announcement, respond with "yes" if the teams can collaborate and "no" if they cannot.

Input:

The first line contains an integer t denoting the number of test cases.
For each test case:
The first line contains two integers n and q, where n is the number of teams and q is the number of queries.

Output:

For each second type of query, print the total number of "yes" and the total number of "no".

Constraints:

1 <= t <= 10
1 <= n <= 10^5
1 <= q <= 10^5

Sample Input:

1
6 6
J 1 3
? 2 3
J 3 6
? 5 1
J 1 5
? 1 3

Sample Output:

1 2

Explanation:

There are 6 teams in the Crypta contest, and a total of 6 announcements are made:
J 1 3: Team 1 and Team 3 begin collaborating.
? 2 3: The answer is NO.
J 3 6: Team 3 and Team 6 start collaborating.
? 5 1: The answer is NO.
J 1 5: Team 1 and Team 5 form a collaboration.
? 1 3: The answer is YES.
*/

import java.util.*;

public class CollabPartnership {

    static class UnionFind {

        private int[] parent;
        private int[] rank;

        public UnionFind(int n) {
            parent = new int[n + 1];
            rank = new int[n + 1];
            for (int i = 1; i <= n; i++) {
                parent[i] = i;
                rank[i] = 0;
            }
        }

        public int find(int x) {
            if (parent[x] != x) {
                parent[x] = find(parent[x]); // Path compression
            }
            return parent[x];
        }

        public void union(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);

            if (rootX != rootY) {
                // Union by rank
                if (rank[rootX] < rank[rootY]) {
                    parent[rootX] = rootY;
                } else if (rank[rootX] > rank[rootY]) {
                    parent[rootY] = rootX;
                } else {
                    parent[rootY] = rootX;
                    rank[rootX]++;
                }
            }
        }

        public boolean isConnected(int x, int y) {
            return find(x) == find(y);
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int t = scanner.nextInt();

        while (t-- > 0) {
            int n = scanner.nextInt();
            int q = scanner.nextInt();

            UnionFind uf = new UnionFind(n);
            int yesCount = 0;
            int noCount = 0;

            for (int i = 0; i < q; i++) {
                String operation = scanner.next();
                int x = scanner.nextInt();
                int y = scanner.nextInt();

                if (operation.equals("J")) {
                    // Join teams x and y
                    uf.union(x, y);
                } else if (operation.equals("?")) {
                    // Query if teams x and y can collaborate
                    if (uf.isConnected(x, y)) {
                        yesCount++;
                    } else {
                        noCount++;
                    }
                }
            }

            System.out.println(yesCount + " " + noCount);
        }

        scanner.close();
    }
}
