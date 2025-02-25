package main

import "fmt"

func main() {
	dp := map[int]int{}
	fmt.Println(getFibonacci(5, dp))
}

func getFibonacci(n int, dp map[int]int) int {
	if n == 1 || n == 0 {
		return n
	}
	if v, ok := dp[n]; ok {
		return v
	}

	dp[n] = getFibonacci(n-1, dp) + getFibonacci(n-2, dp)
	return dp[n]
}
