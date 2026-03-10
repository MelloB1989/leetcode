package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	stack := []int{}
	n := len(nums)
	for i := range result {
		result[i] = -1
	}
	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = nums[i%n]
		}
		stack = append(stack, i%n)
	}
	return result
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
