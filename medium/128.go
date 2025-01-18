package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	count := 0
	max := 0
	// fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		// fmt.Println(count, nums[i], nums[i]-nums[i-1])
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			count += 1
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
		count = 0
	}
	return max + 1
}

func main() {
	fmt.Println(longestConsecutive([]int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}))
}
