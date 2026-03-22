package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	stack := []rune{}
	for _, d := range num {
		for len(stack) > 0 && d < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, d)
	}
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}
	stack = stack[i:]

	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

func main() {
	fmt.Println(removeKdigits("10200", 1))
}
