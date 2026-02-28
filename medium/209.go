package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	least := len(nums) + 1
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < least {
				least = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if least == len(nums)+1 {
		return 0
	}
	return least
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
