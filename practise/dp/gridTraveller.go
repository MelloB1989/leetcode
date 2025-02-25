package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 0, 1},
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 1, 1},
	}

	m, n := len(grid), len(grid[0])
	dp := map[string]int{}
	fmt.Println(travel(m-1, n-1, grid, dp))
}

func travel(m, n int, grid [][]int, dp map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)

	if v, ok := dp[key]; ok {
		return v
	}

	if m < 0 || n < 0 || grid[m][n] == 0 {
		return 0
	}

	if m == 0 && n == 0 {
		return 1
	}

	dp[key] = travel(m-1, n, grid, dp) + travel(m, n-1, grid, dp)
	return dp[key]
}
