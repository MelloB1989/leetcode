package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	write := 1
	uc := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			uc += 1
			if uc <= 2 {
				nums[write] = nums[i]
				write += 1
			}
		} else {
			uc = 1
			nums[write] = nums[i]
			write += 1
		}
	}

	return write
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 4, 4}))
	//0 0 1 1 2 3 3 4 4 .......
}
