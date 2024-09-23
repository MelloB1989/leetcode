package main

import "fmt"

func lengthOfLastWord(s string) int {
	start := 0
	end := len(s) - 1
	for end >= start && s[end] == ' ' {
		end--
	}
	s = s[start : end+1]
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return len(string(s[i : len(s)-1]))
		}
	}
	return len(s)
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
