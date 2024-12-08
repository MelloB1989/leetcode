package main

import "fmt"

/*
[24,12,8,6]
 1  0  0 0
 2  2  1 1
 3  3  3 2
*/

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for j := 0; j < len(nums); j++ {
		result[j] = 1
	}

	left := 1
	right := 1

	for i := 0; i < len(nums); i++ {
		result[i] *= left
		left *= nums[i]
	}

	for j := len(nums) - 1; j >= 0; j-- {
		result[j] *= right
		right *= nums[j]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
