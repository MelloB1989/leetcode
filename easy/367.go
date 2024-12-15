package main

import "fmt"

func isPerfectSquare(num int) bool {
	n, x := num, num
	y := 1
	e := 0.1
	for float64(x-y) > e {
		x = (x + y) / 2
		y = n / x
	}
	if int(x)*int(x) == num {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(6))
}
