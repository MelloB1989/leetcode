package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	r := len(nums)
	for i := 0; i < len(nums); i++ {
		r += i - nums[i]
	}
	return r
}

func main() {
	// fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{3, 4, 0, 1}))
}
