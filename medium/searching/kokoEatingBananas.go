package main

import (
	"fmt"
)

func minEatingSpeed(piles []int, h int) int {

	max := piles[0]
	for _, v := range piles {
		if max < v {
			max = v
		}
	}

	var canEat func(k int) bool
	canEat = func(k int) bool {
		hourSpent := 0
		for _, v := range piles {
			hourSpent += (v + k - 1) / k
			if hourSpent > h {
				return false
			}
		}
		if hourSpent <= h {
			return true
		}
		return false
	}
	left := 1
	for left < max {
		mid := (left + max) / 2
		fmt.Println(mid, canEat(mid))
		if canEat(mid) {
			max = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Print(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}
