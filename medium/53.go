package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
