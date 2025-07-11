/*
Building Layout Exploration (Layoutexplore)
Building Layout Exploration

Problem Description

You are tasked with analyzing the structure of a building represented by an n × m grid. The building has open rooms and walls, and your objective is to determine the count of distinct rooms within the layout. The grid comprises floor squares ('.') and wall squares ('#'). You have the freedom to move left, right, up, and down through the floor squares.

Input

The first line contains two integers, n and m, separated by a space (1 ≤ n, m ≤ 1000), indicating the height and width of the building.
Following are n lines of m characters describing the building's layout. Each character is either '.' (indicating a floor) or '#' (indicating a wall).

Output

Print a single integer: the total number of distinct rooms in the building.

Constraints
1 ≤ n, m ≤ 1000

Example

5 8
########
#..#...#
####.#.#
#..#...#
########

Output

3

Note

In this scenario, the building contains three distinct rooms, each surrounded by walls ('#').
*/
import java.util.*;

public class BuildingLayoutExploration {

    // Direction arrays for moving up, right, down, left
    private static final int[] dx = { -1, 0, 1, 0 };
    private static final int[] dy = { 0, 1, 0, -1 };

    public static int countRooms(char[][] grid) {
        if (grid == null || grid.length == 0) return 0;

        int n = grid.length;
        int m = grid[0].length;
        int roomCount = 0;

        // Create a visited array to keep track of explored cells
        boolean[][] visited = new boolean[n][m];

        // Iterate through each cell in the grid
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                // If we find an unvisited floor cell, explore it
                if (grid[i][j] == '.' && !visited[i][j]) {
                    dfs(grid, i, j, visited);
                    roomCount++;
                }
            }
        }

        return roomCount;
    }

    private static void dfs(char[][] grid, int x, int y, boolean[][] visited) {
        // Mark current cell as visited
        visited[x][y] = true;

        // Explore all four directions
        for (int i = 0; i < 4; i++) {
            int newX = x + dx[i];
            int newY = y + dy[i];

            // Check if the new position is valid and is an unvisited floor
            if (
                isValid(grid, newX, newY) &&
                grid[newX][newY] == '.' &&
                !visited[newX][newY]
            ) {
                dfs(grid, newX, newY, visited);
            }
        }
    }

    private static boolean isValid(char[][] grid, int x, int y) {
        return x >= 0 && x < grid.length && y >= 0 && y < grid[0].length;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        // Read dimensions
        int n = scanner.nextInt();
        int m = scanner.nextInt();
        scanner.nextLine(); // Consume the newline

        // Create and fill the grid
        char[][] grid = new char[n][m];
        for (int i = 0; i < n; i++) {
            String line = scanner.nextLine();
            grid[i] = line.toCharArray();
        }

        // Calculate and print the result
        int result = BuildingLayoutExploration.countRooms(grid);
        System.out.println(result);

        scanner.close();
    }
}
