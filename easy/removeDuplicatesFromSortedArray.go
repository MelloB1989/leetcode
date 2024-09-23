package main

import "fmt"

func removeDuplicates(nums []int) int {
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
