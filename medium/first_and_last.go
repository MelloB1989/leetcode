package main

import "fmt"

func searchRange(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}
	r := []int{-1, -1}
	if len(result) > 0 {
		r = []int{result[0], result[len(result)-1]}
	}
	return r
}

// func searchRange(nums []int, target int) []int {
// 	var result []int
// 	left := 0
// 	right := len(nums) - 1
// 	mid := int((right + left) / 2)
// 	fmt.Print(mid)
// 	for left < right {
// 		// fmt.Print(left, mid, right)
// 		if nums[mid] == target {
// 			result = append(result, mid)
// 		}
// 		if target < nums[mid] {
// 			right = mid
// 			mid = int((right + left) / 2)
// 		} else {
// 			left = mid
// 			mid = int((right + left) / 2)
// 		}
// 	}
// 	r := []int{-1, -1}
// 	if len(result) > 0 {
// 		r = []int{result[0], result[len(result)-1]}
// 	}
// 	return r
// }

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
