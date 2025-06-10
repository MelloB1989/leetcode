package main

import "fmt"

/*
Bounds:

rows: 1 to m
cols: 1 to n

given at (i, j) I have 4 moves: up, down, right and left
since my finish is in bottom right, i need to go in down and left direction only
so at (i, j) I have 2 moves: down and right
so the bot can arrive from: (i-1,j) for (i, j-1) (up or left)

my recurrence:
    dp[i][j] = dp[i-1][j] + dp[i][j-1]
base case:
    dp[0][j] = 1
    dp[i][0] = 1

*/

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	fmt.Println(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			fmt.Println(dp)
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
}
