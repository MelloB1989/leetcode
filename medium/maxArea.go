package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	if r < 2 {
		return min(height[r], height[l])
	}
	max := 0
	for r > l {
		width := r - l
		area := min(height[r], height[l]) * width
		// fmt.Println("checking for h: ", min(height[r], height[l]), " w: ", width, " l: ", l, " r: ", r, " area: ", area)
		if area > max {
			max = area
		}
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return max
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
}
