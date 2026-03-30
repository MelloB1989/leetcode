package main

import (
	"fmt"
)

func maxWidthRamp(nums []int) int {
	stack := []int{}

	for i, n := range nums {
		if len(stack) == 0 || n < nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	maxWidth := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			maxWidth = max(maxWidth, i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return maxWidth
}

func main() {
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
