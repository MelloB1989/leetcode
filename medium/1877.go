package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	max := 0
	l, r := 0, n-1

	for l < r {
		if nums[l]+nums[r] > max {
			max = nums[l] + nums[r]
		}
		l++
		r--
	}

	return max
}

func main() {
	fmt.Println(minPairSum([]int{3, 5, 2, 3}))
}
