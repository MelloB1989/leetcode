package main

import "fmt"

func maximumRemovals(s string, p string, removable []int) int {
	max := len(removable)
	isSubsequence := func(s string, p string, k int) bool {
		removed := make([]bool, len(s))
		for i := 0; i < k; i++ {
			removed[removable[i]] = true
		}

		j := 0
		for i := 0; i < len(s) && j < len(p); i++ {
			if !removed[i] && s[i] == p[j] {
				j++
			}
		}
		return j == len(p)
	}
	left := 0
	for left < max {
		mid := (left + max + 1) / 2
		if isSubsequence(s, p, mid) {
			left = mid
		} else {
			max = mid - 1
		}
	}
	return left
}

func main() {
	// fmt.Println(maximumRemovals("abcacb", "ab", []int{3, 1, 0}))
	fmt.Println(maximumRemovals("abcbddddd", "abcd", []int{3, 2, 1, 4, 5, 6}))
}
