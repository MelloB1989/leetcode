package main

import (
	"fmt"
)

// func findApples(n int) int {
// 	var calApples func(sideLen int) int
// 	calApples = func(sideLen int) int {
// 		totalApples := 0
// 		halfSide := sideLen / 2

// 		var minX, maxX int
// 		if sideLen%2 == 0 {
// 			// Even side length
// 			minX = -halfSide
// 			maxX = halfSide - 1
// 		} else {
// 			// Odd side length
// 			minX = -halfSide
// 			maxX = halfSide
// 		}

// 		for x := minX; x <= maxX; x++ {
// 			for y := minX; y <= maxX; y++ {
// 				apples := int(math.Abs(float64(x)) + math.Abs(float64(y)))
// 				totalApples += apples
// 			}
// 		}
// 		return totalApples
// 	}

// 	// calApples = func(sideLen int) int {
// 	// 	totalApples := 0
// 	// 	for x := 0; x < sideLen; x++ {
// 	// 		for y := 0; y < sideLen; y++ {
// 	// 			totalApples += x + y
// 	// 		}
// 	// 	}
// 	// 	return totalApples * 3
// 	// }

// 	left := 0
// 	right := 2
// 	for calApples(right) < n {
// 		right *= 2
// 	}
// 	for left < right {
// 		mid := (left + right) / 2
// 		if calApples(mid) < n {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
// 	return left * 4
// }

func findApples(n int) int {
	var S func(k int64) int64
	S = func(k int64) int64 {
		return 4*k*k*k + 6*k*k + 2*k
	}

	right := int64(1)
	left := int64(0)

	for S(right) < int64(n) {
		right *= 2
	}

	for left < right {
		mid := left + (right-left)/2
		if S(mid) >= int64(n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int(8 * left)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println("Apples", findApples(9980017896001))
}
