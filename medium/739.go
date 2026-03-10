package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	answers := make([]int, len(temperatures))
	stack := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 {
			j := stack[len(stack)-1]
			if temperatures[i] < temperatures[j] {
				answers[i] = j - i
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}
	return answers
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
