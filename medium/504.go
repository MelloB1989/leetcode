package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] != nums[mid+1] {
			right = mid
		} else {
			left = mid + 2
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
