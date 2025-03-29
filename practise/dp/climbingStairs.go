package main

import "fmt"

func main() {
	dp := map[int]int{}
	n := 30
	fmt.Println(getStairs(n, dp))
	fmt.Println(dp)
}

func getStairs(n int, dp map[int]int) int {
	if v, ok := dp[n]; ok {
		return v
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	dp[n] = getStairs(n-1, dp) + getStairs(n-2, dp)
	return dp[n]
}
