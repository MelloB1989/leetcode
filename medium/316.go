package main

import "fmt"

func removeDuplicateLetters(s string) string {
	result := ""
	stack := []rune{}
	count := make(map[rune]int)
	seen := make(map[rune]bool)
	for _, c := range s {
		count[c]++
	}

	for _, c := range s {
		count[c]--
		if seen[c] {
			continue
		}
		for len(stack) > 0 &&
			count[stack[len(stack)-1]] > 0 &&
			stack[len(stack)-1] > c {
			seen[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		seen[c] = true
	}

	for _, i := range stack {
		result += string(i)
	}

	return result
}

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
