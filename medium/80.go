package main

import "fmt"

func search(nums []int, target int) bool {
	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		//Skip duplicates
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] { //Lies on Left side
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	// fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 1))
}
