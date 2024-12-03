package main

import (
	"fmt"
)

func findTargetSumWays(nums []int, target int) int {
	count := 0
	var find func(i, sum int)
	find = func(i, sum int) {
		if i >= len(nums) {
			if sum == target {
				count += 1
			}
			return
		}
		find((i + 1), sum-nums[i])
		find((i + 1), sum+nums[i])
	}
	find(0, 0)
	return count
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
