package main

import "fmt"

func trap(height []int) int {
	stack := []int{}
	water := 0
	for i, n := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < n {
			bottom := stack[len(stack)-1]
			right := i
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			h := min(height[left], height[right]) - height[bottom]
			width := right - left - 1
			water += h * width
		}
		stack = append(stack, i)
	}
	return water
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
