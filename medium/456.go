package main

import "fmt"

func find132pattern(nums []int) bool {
	stack := []int{}
	second := -1 << 31
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < second {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			second = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
}
