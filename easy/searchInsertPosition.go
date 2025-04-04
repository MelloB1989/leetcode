package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = int((left + right) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
