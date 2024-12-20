package main

import "fmt"

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
