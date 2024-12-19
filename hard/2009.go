package main

import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	unique := nums[:j+1]
	minOps := n
	right := 0
	for left, _ := range unique {
		for right < len(unique) && unique[right] < unique[left]+n {
			right += 1
		}
		minOps = min(minOps, n-(right-left))
	}
	return minOps
}

func main() {
	fmt.Println(minOperations([]int{8, 5, 9, 9, 8, 4}))
}
