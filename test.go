package main

import "fmt"

func main() {
	a := []int{6, 18, 19, 7, 16}
	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		if i < len(a)-1 && a[i] < a[i+1] {
			b[i] = a[i+1]
		} else {
			b[i] = -1
		}
	}

	fmt.Println(b)
}
