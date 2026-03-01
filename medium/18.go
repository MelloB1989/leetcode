package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	for a, _ := range nums {
		if a == n-3 {
			break
		}
		// Skip dups
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		b := a + 1

		for b < n-2 {
			// Skip dups
			if b > a+1 && nums[b] == nums[b-1] {
				b++
				continue
			}

			l, r := b+1, n-1

			for l < r {
				sum := nums[a] + nums[b] + nums[l] + nums[r]

				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					result = append(result, []int{nums[a], nums[b], nums[l], nums[r]})

					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}

					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
			b++
		}
	}
	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	// nums := []int{2, 2, 2, 2, 2}
	target := 0

	fmt.Println(fourSum(nums, target))
}
