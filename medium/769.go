package main

import (
	"fmt"
)

func maxChunksToSorted(arr []int) int {
	stack := []int{}
	for _, n := range arr {
		curMax := n
		for len(stack) > 0 && stack[len(stack)-1] > n {
			curMax = max(curMax, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, curMax)
	}
	return len(stack)
}

func main() {
	// fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4})) // 4, 3, 2, 1, 0
	fmt.Println(maxChunksToSorted([]int{1, 2, 0, 3}))
}
