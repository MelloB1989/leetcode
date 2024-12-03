package main

import (
	"fmt"
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	n := int(math.Pow(2, float64(len(nums))))
	count := 0
	for i := 0; i < n; i++ {
		signs := make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			if (i>>j)&1 == 1 {
				signs[j] = -nums[j]
			} else {
				signs[j] = nums[j]
			}
		}
		sum := 0
		for p := 0; p < len(signs); p++ {
			sum += signs[p]
		}
		if sum == target {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(findTargetSumWays([]int{6, 44, 30, 25, 8, 26, 34, 22, 10, 18, 34, 8, 0, 32, 13, 48, 29, 41, 16, 30}, 12))
}
