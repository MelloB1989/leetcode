package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	var b func(current string, open int, close int)
	b = func(current string, open int, close int) {
		if len(current) == n*2 {
			result = append(result, current)
			return
		}
		if open < n {
			b(current+"(", open+1, close)
		}
		if close < open {
			b(current+")", open, close+1)
		}
	}
	b("", 0, 0)
	return result
}

func main() {
	fmt.Println(generateParenthesis(1))
}
