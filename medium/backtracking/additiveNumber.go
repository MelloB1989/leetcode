package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	n := len(num)
	var check func(p1, p2, start int) bool
	check = func(p1, p2, start int) bool {
		if start == n {
			return true
		}

		sum := strconv.Itoa(p1 + p2)
		fmt.Println(p1, p2, sum)

		if start+len(sum) > n || num[start:start+len(sum)] != sum {
			return false
		}

		return check(p2, p1+p2, start+len(sum))
	}
	for i := 1; i < n; i++ {
		if num[0] == '0' && i > 1 {
			break
		}
		p, _ := strconv.Atoi(num[:i])
		for j := i + 1; j < n; j++ {
			if num[i] == '0' && j > i+1 {
				break
			}
			q, _ := strconv.Atoi(num[i:j])
			// fmt.Println(p, q, j)
			if check(p, q, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(isAdditiveNumber("123"))
}
