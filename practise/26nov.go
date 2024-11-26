package main

import (
	"fmt"
)

// max and second max
func main() {
	arr := []int{9, 7, 8}
	m := arr[0]
	s := arr[1]
	if arr[1] > m {
		m = arr[1]
		s = arr[0]
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] > m {
			s = m
			m = arr[i]
		} else if arr[i] > s {
			s = arr[i]
		}
	}
	fmt.Print(m, s)
}
