package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSum := map[int]int{}
	count := 0
	preSum[0] = 1
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if f, exists := preSum[sum-k]; exists {
			count += f
		}

		preSum[sum] += 1
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
