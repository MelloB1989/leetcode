package main

import "fmt"

func majorityElement(nums []int) int {
	majority := 0
	res := 0
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
		if m[nums[i]] > majority {
			res = nums[i]
			majority = m[res]
		}
	}
	// fmt.Println(m, res, majority)
	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3, 4, 4, 3}))
}
