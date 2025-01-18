package main

import (
	"fmt"
	"sort"
)

func maximumGap(nums []int) int {
	s := len(nums)
	if s < 2 {
		return 0
	}
	sort.Ints(nums)
	fmt.Println(nums)
	maxDiff := 0
	for i := 1; i < s; i++ {
		diff := nums[i] - nums[i-1]
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
}
