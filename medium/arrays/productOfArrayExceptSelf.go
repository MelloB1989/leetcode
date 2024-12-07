package main

import "fmt"

func productExceptSelf(nums []int) []int {
	pr := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			pr *= 1
		} else {
			pr *= nums[i]
		}
	}

	fmt.Println(pr)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = pr
		} else {
			nums[i] = pr / nums[i]
		}
	}

	return nums
}

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
