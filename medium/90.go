package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	set := []int{}

	sort.Ints(nums)

	var b func(start int)
	b = func(start int) {
		s := make([]int, len(set))
		copy(s, set)
		result = append(result, s)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			set = append(set, nums[i])
			b(i + 1)
			set = set[:len(set)-1]
		}
	}
	b(0)
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
