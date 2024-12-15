package main

import "fmt"

func search(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		fmt.Println(nums[mid])
		if nums[left] <= nums[mid] {
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
	return -1
}

func main() {
	fmt.Println(search([]int{5, 1, 3}, 3))
	// fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 1))
}
