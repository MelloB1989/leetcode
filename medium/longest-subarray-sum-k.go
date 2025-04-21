// Input: arr[] = [10, 5, 2, 7, 1, -10], k = 15
// Output: 6
// Explanation: Subarrays with sum = 15 are [5, 2, 7, 1], [10, 5] and [10, 5, 2, 7, 1, -10]. The length of the longest subarray with a sum of 15 is 6.

// Input: arr[] = [-5, 8, -14, 2, 4, 12], k = -5
// Output: 5
// Explanation: Only subarray with sum = 15 is [-5, 8, -14, 2, 4] of length 5.

// Input: arr[] = [10, -10, 20, 30], k = 5
// Output: 0
// Explanation: No subarray with sum = 5 is present in arr[].

package main

import "fmt"

func longestSubarrayWithSumK(nums []int, k int) int {
	oc := map[int]int{}
	oc[0] = -1
	maxLen := 0
	prefix := 0
	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		if v, exists := oc[prefix-k]; exists {
			l := i - v
			if l > maxLen {
				maxLen = l
			}
		} else {
			oc[prefix-k] = i
		}
	}
	return maxLen
}

func main() {
	nums := []int{-5, 8, -14, 2, 4, 12}
	k := -5
	fmt.Println(longestSubarrayWithSumK(nums, k))
}
