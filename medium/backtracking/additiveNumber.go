package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	result := true
	var b func(p1s int, p1e int, p2s int, p2e int) bool
	b = func(p1s int, p1e int, p2s int, p2e int) bool {
		if p1e < len(num) && p2e < len(num) {
			p, _ := strconv.Atoi(num[p1s:p1e])
			q, _ := strconv.Atoi(num[p2s:p2e])
			next := strconv.Itoa((p + q))
			if p2e+len(next) >= len(num) {
				result = false
				return result
			}
			fmt.Println(num[p1s:p1e], num[p2s:p2e], num[p2e:p2e+len(next)], next)
			if next != num[p2e:p2e+len(next)] {
				result = b(p1s, p1e, p2s, p2e+1) || b(p1s, p1e+1, p2s+1, p2e+1)
				return result
			} else {
				result = true
				b(p1s+1, p1e+1, p2s+1, p2e+1)
				return result
			}
		}
		return result
	}
	return b(0, 1, 1, 2)
}

func main() {
	fmt.Println(isAdditiveNumber("199100199"))
}
