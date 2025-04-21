package main

import (
	"fmt"
)

func maximumPossibleSize(nums []int) int {
	c := 0
	locMax := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= locMax {
			locMax = nums[i]
			c++
		}
	}
	return c
}

func main() {
	nums := []int{19, 80, 63, 74}
	fmt.Println(maximumPossibleSize(nums))
}
