package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	result := make([]int, len(nums1))
	nge := map[int]int{}

	for i := range nums2 {
		for len(stack) > 0 &&
			nums2[stack[len(stack)-1]] < nums2[i] {

			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			nge[nums2[idx]] = nums2[i]
		}

		stack = append(stack, i)
	}

	for i, num := range nums1 {
		if v, ok := nge[num]; ok {
			result[i] = v
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
