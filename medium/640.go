package main

import (
	"fmt"
	"regexp"
	"strings"
)

func splitWithOperators(exp string) []string {
	re := regexp.MustCompile(`[-+]`)
	indices := re.FindAllStringIndex(exp, -1)

	if len(indices) == 0 {
		return []string{exp}
	}

	var tokens []string
	lastIndex := 0

	if indices[0][0] == 0 {
		lastIndex = 0
		indices = indices[1:] // Skip the first operator
	}

	for _, idx := range indices {
		if lastIndex < idx[0] {
			tokens = append(tokens, exp[lastIndex:idx[0]])
		}
		tokens = append(tokens, exp[idx[0]:idx[1]+1])
		lastIndex = idx[1]
	}

	if lastIndex < len(exp) {
		tokens = append(tokens, exp[lastIndex:])
	}

	return tokens
}

func solveEquation(equation string) string {
	var eval func(eq string) []int
	eval = func(eq string) []int {
		tokens := splitWithOperators(eq)
		res := make([]int, 2)
		for _, token := range tokens {
			if token ==
		}
		return []int{}
	}
	eval(strings.Split(equation, "=")[0])
	return ""
}

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
}
