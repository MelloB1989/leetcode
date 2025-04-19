package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := [][]int{}
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, length-1
			for l < r {
				a, b, c, d := nums[i], nums[j], nums[l], nums[r]
				sum := a + b + c + d
				if sum == target {
					result = append(result, []int{a, b, c, d})

					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return result
}

func main() {
	// nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{2, 2, 2, 2, 2}
	target := 8

	fmt.Println(fourSum(nums, target))
}
