package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	start, end := -1, -1
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			start = i
			break
		}
	}

	if start == -1 {
		return 0
	}

	for i := n - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			end = i
			break
		}
	}

	minVal, maxVal := nums[start], nums[start]
	for i := start; i <= end; i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	for start > 0 && nums[start-1] > minVal {
		start--
	}

	for end < n-1 && nums[end+1] < maxVal {
		end++
	}

	return end - start + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
}
