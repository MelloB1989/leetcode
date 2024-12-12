package main

import "fmt"

func isBadVersion(version int) bool {
	if version >= 6 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := (left + right) / 2
		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(firstBadVersion(8))
}
