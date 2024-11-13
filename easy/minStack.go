package main

import "fmt"

func main() {
	a := []int{19, 6, 20, 1, 21}
	b := make([]int, len(a))
	min := a[0]

	for i := 0; i < len(a); i++ {
		if a[i] <= min {
			min = a[i]
		}
		b[i] = min
	}

	fmt.Println(b)
}
