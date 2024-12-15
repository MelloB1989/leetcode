package main

import "fmt"

func splitArray(nums []int, k int) int {
	n := len(nums)
	max := nums[0]
	sum := 0

	for i := 0; i < n; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		sum += nums[i]
	}
	var hasKSplits func(sum int) bool
	hasKSplits = func(sum int) bool {
		splits := 0
		s := 0
		for i := 0; i < n; i++ {
			// fmt.Println("Current sum:", s, "Current element:", preSum[i], "Allowed sum:", sum, "Splits:", splits)
			if s+nums[i] > sum {
				splits++
				s = 0
			}
			s += nums[i]
		}
		return splits+1 <= k
	}
	left, right := max, sum
	for left < right {
		mid := (left + right) / 2
		fmt.Println(left, right, mid)
		if hasKSplits(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(splitArray([]int{1, 4, 4}, 3))
}
