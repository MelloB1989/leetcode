package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	min := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if nums[left] < min {
			min = nums[left]
		}
		if nums[mid] < min {
			min = nums[mid]
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return min
}

func main() {
	fmt.Println(findMin([]int{2, 0, 1, 1, 1}))
}
