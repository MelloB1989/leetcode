package main

import "fmt"

func missingNumber(nums []int) int {
	missing := 0
	a := 0
	b := 0
	for i := 0; i <= len(nums); i++ {
		if i < len(nums) {
			a += nums[i]
		}
		b += i
	}
	missing = b - a
	return missing
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	// fmt.Println(missingNumber([]int{3,0,1}))
}
