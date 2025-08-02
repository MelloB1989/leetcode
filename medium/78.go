package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}
	set := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		s := make([]int, len(set))
		copy(s, set)
		result = append(result, s)
		for i := start; i < len(nums); i++ {
			set = append(set, nums[i])
			backtrack(i + 1)
			set = set[:len(set)-1]
		}
	}
	backtrack(0)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
