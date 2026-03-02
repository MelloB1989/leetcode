package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))

	fmt.Println(l, r)

	for l <= r {
		sum := l*l + r*r
		if sum > c {
			r--
		} else if sum < c {
			l++
		} else {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(judgeSquareSum(4))
}
