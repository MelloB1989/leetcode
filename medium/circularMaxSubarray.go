package main

import "fmt"

func kadane(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubarraySumCircular(nums []int) int {
	s := kadane(nums)
	tot := 0
	for i := 0; i < len(nums); i++ {
		tot += nums[i]
		nums[i] = -nums[i]
	}

	tot = tot + kadane(nums)
	if tot > s && tot != 0 {
		return tot
	}
	return s
}

func main() {
	fmt.Print(maxSubarraySumCircular([]int{-3, -2, -3}))
}
