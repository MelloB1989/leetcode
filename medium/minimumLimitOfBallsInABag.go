package main

import (
	"fmt"
	"sort"
)

type IntSliceDesc []int

func (s IntSliceDesc) Len() int {
	return len(s)
}

func (s IntSliceDesc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSliceDesc) Less(i, j int) bool {
	return s[i] > s[j]
}

func minimumSize(nums []int, maxOperations int) int {
	splitMap := map[int][]int{}

	sort.Sort(IntSliceDesc(nums))

	for i := 0; i < len(nums); i++ {
		splitMap[nums[i]] =
	}

	var split func()

	for i := 0; i < maxOperations; i++ {

	}
}

func main() {
	fmt.Println(minimumSize([]int{10, 15, 20}, 6))
}
