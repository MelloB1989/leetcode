package main

import "fmt"

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	curr := 0
	prev1 := 3
	prev2 := 2

	for i := 3; i < n; i++ {
		curr = prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return curr
}

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}
