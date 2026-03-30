package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	for i := 0; i <= len(heights); i++ {
		var h int
		if i == len(heights) {
			h = 0
		} else {
			h = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[top]

			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack)-1] - 1
			}

			a := height * w
			if maxArea < a {
				maxArea = a
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{4, 2}))
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
