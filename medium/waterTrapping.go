package main

import "fmt"

// Water Trapping
func main() {
	heights := []int{7, 4, 0, 9}
	l := 0
	r := len(heights) - 1
	w := 0
	lM, rM := heights[l], heights[r]

	for r >= l {
		if lM < rM {
			if heights[l] < lM {
				w += lM - heights[l]
			} else {
				lM = heights[l]
			}
			l++
		} else {
			if rM > heights[r] {
				w += rM - heights[r]
			} else {
				rM = heights[r]
			}
			r--
		}
	}

	fmt.Println(w)
}
