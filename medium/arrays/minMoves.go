package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	median := nums[n/2]
	moves := 0
	for _, num := range nums {
		moves += abs(num - median)
	}
	return moves
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(minMoves2([]int{1, 2, 3}))
}
