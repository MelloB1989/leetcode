package main

import (
	"fmt"
)

//	func fib(n int) int {
//		if n == 0 || n == 1 {
//			return n
//		}
//		return fib(n-1) + fib(n-2)
//	}
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}

func main() {
	fmt.Println(fib(4))
}
