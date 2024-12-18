package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	num1Map := make(map[int]int)
	num2Map := make(map[int]int)
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	for i := 0; i < len(nums1); i++ {
		num1Map[nums1[i]] += 1
		if i < len(nums2) {
			num2Map[nums2[i]] += 1
		}
	}
	for key, _ := range num1Map {
		if num2Map[key] >= 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
