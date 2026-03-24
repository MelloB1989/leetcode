package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	stack := []int{}

	for i, _ := range nums {
		if i < len(nums)-1 && nums[i] > nums[i+1] {
			stack = append(stack, i)
		}
	}
	if len(stack) < 2 {
		if len(stack) == 0 {
			return 0
		} else {
			return len(nums)
		}
	}
	i, j := stack[0], stack[len(stack)-1]

	return j - i + 2
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
}
