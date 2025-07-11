/*
Snakes and Ladders
Snakes and Ladders

In the Snakes and Ladders game, Markov, a player, wonders how many moves he needs to make to reach square 100, taking advantage of the ability to roll the die to any desired number. The game's rules are standard:

A cubic die with 6 faces, numbered 1 to 6, is used.
Players start from square 1 and aim to reach square 100 with an exact roll of the die. If the roll would move them beyond square 100, they stay in place.
Ladders allow players to climb up. Ladders only go upwards.
Snakes force players to move down the snake and exit through the tail. Snakes only go downwards.
Function Description

Implement the quickestWayUp function, which should return an integer representing the fewest number of moves required.

The function has the following parameters:

ladders: A 2D integer array where each ladders[i] contains the start and end cell numbers of a ladder.
snakes: A 2D integer array where each snakes[i] contains the start and end cell numbers of a snake.

Input Format

The first line contains the number of test cases, t.

For each test case:

The first line contains n, the number of ladders.
Each of the next n lines contains two space-separated integers, denoting the start and end of a ladder.
The next line contains the integer m, the number of snakes.
Each of the next m lines contains two space-separated integers, indicating the start and end of a snake.

Constraints

1 <= t <= 10
1 <= n, m <= 15
The game board is always 10x10, with squares numbered from 1 to 100.
Square 1 or square 100 will not be the starting points of any ladder or snake.
Each square will have at most one endpoint from either a snake or a ladder.

Output Format

For each test case, print the minimum number of rolls needed to move from the start to the finish on a separate line. If there's no solution, print -1.

Sample Input

2
3
32 62
42 68
12 98
7
95 13
97 25
93 37
79 27
75 19
49 47
67 17
4
8 52
6 80
26 42
2 72
9
51 19
39 11
37 29
81 3
59 5
79 23
53 7
43 33
77 21

Sample Output

3
5

Explanation

For the first test:

Markov rolls a 5 and then a 6, landing at square 12. There's a ladder to square 98. A roll of 2 completes the traverse in 3 rolls.

For the second test:

Markov initially rolls a 5, climbing the ladder to square 80. Three rolls of 6 get to square 98. A final roll of 2 lands on the target square in 5 total rolls
*/
import java.io.*;
import java.util.*;

public class SnakesAndLadder {

    public static int quickestWayUp(int[][] ladders, int[][] snakes) {
        // Create a map to store ladder and snake endpoints
        Map<Integer, Integer> board = new HashMap<>();

        // Add ladders to the board
        for (int[] ladder : ladders) {
            board.put(ladder[0], ladder[1]);
        }

        // Add snakes to the board
        for (int[] snake : snakes) {
            board.put(snake[0], snake[1]);
        }

        // BFS to find minimum moves
        Queue<int[]> queue = new LinkedList<>();
        boolean[] visited = new boolean[101]; // squares 1-100

        queue.offer(new int[] { 1, 0 }); // {position, moves}
        visited[1] = true;

        while (!queue.isEmpty()) {
            int[] current = queue.poll();
            int position = current[0];
            int moves = current[1];

            // Try all possible dice rolls (1-6)
            for (int dice = 1; dice <= 6; dice++) {
                int nextPosition = position + dice;

                // If we would go beyond 100, stay in place
                if (nextPosition > 100) {
                    continue;
                }

                // Check if there's a ladder or snake at this position
                if (board.containsKey(nextPosition)) {
                    nextPosition = board.get(nextPosition);
                }

                // If we reached square 100, return the number of moves
                if (nextPosition == 100) {
                    return moves + 1;
                }

                // If we haven't visited this position, add it to queue
                if (!visited[nextPosition]) {
                    visited[nextPosition] = true;
                    queue.offer(new int[] { nextPosition, moves + 1 });
                }
            }
        }

        // If we can't reach square 100
        return -1;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int t = scanner.nextInt();

        for (int testCase = 0; testCase < t; testCase++) {
            // Read number of ladders
            int n = scanner.nextInt();
            int[][] ladders = new int[n][2];

            for (int i = 0; i < n; i++) {
                ladders[i][0] = scanner.nextInt();
                ladders[i][1] = scanner.nextInt();
            }

            // Read number of snakes
            int m = scanner.nextInt();
            int[][] snakes = new int[m][2];

            for (int i = 0; i < m; i++) {
                snakes[i][0] = scanner.nextInt();
                snakes[i][1] = scanner.nextInt();
            }

            // Find and print the result
            int result = quickestWayUp(ladders, snakes);
            System.out.println(result);
        }

        scanner.close();
    }
}
