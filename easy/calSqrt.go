package main

import "fmt"

func mySqrt(x int) int {
	n := x
	y := 1
	e := 0.1
	for float64(x-y) > e {
		x = (x + y) / 2
		y = n / x
	}
	return int(x)
}

func main() {
	x := 31
	fmt.Println(mySqrt(x))
}
