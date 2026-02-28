package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	soln := [][]int{}
	for i, _ := range nums {
		if i == n-2 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				soln = append(soln, []int{nums[i], nums[l], nums[r]})
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
	}
	return soln
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
