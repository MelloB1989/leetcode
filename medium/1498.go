package main

import (
	"fmt"
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := 0

	const mod = 1_000_000_007

	// Precompute powers of 2
	pow := make([]int, n)
	pow[0] = 1
	for i := 1; i < n; i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}

	l, r := 0, n-1

	for l <= r {
		if nums[l]+nums[r] <= target {
			result = (result + pow[r-l]) % mod
			l++
		} else {
			r--
		}
	}

	return result
}

func main() {
	// fmt.Println(numSubseq([]int{3, 5, 6, 7}, 9))
	fmt.Println(numSubseq([]int{7, 10, 7, 3, 7, 5, 4}, 12))
}
