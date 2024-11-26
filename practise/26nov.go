package main

import (
	"fmt"
)

// max and second max
func main() {
	arr := []int{8, 9, 5, 8, 3, 0, 2}
	m := arr[0]
	s := m
	for i := 0; i < len(arr); i++ {
		if arr[i] > m {
			s = m
			m = arr[i]
		}
	}
	fmt.Print(m, s)
}
