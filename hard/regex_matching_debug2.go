package main

import "fmt"

func isMatch(s string, p string) bool {
	pointer := 0
	patterns := parsePattern(p) // Assuming parsePattern returns []map[string]string
	length := len(s)

	for i := 0; i < len(patterns); i++ {
		currentPattern := patterns[i] // Get the current map pattern
		matched := false              // Track if we matched any character for '+' or '*'

		for key, value := range currentPattern {
			switch key {
			case "+":
				// Ensure at least one match for "+"
				if pointer < length && (value == string(s[pointer]) || value == ".") {
					matched = true // We have a match
					pointer += 1   // Consume this character
					// Continue to consume additional characters if they match
					for pointer < length && (value == string(s[pointer]) || value == ".") {
						pointer += 1
					}
				}
			case ".":
				if pointer < length { // Only increment if there's a character to match
					matched = true
					pointer += 1 // Match any single character
				}
			case "*":
				// Check for zero occurrences first
				if pointer < length && (value == string(s[pointer]) || value == ".") {
					// Match one or more occurrences
					for pointer < length && (value == string(s[pointer]) || value == ".") {
						pointer += 1
					}
				} else {
					matched = true // '*' can match zero occurrences
				}
			}
		}

		// Check if we have a '+' that required a match and none was found
		if !matched && (len(currentPattern) == 1 && (currentPattern["+"] == "" || currentPattern["."] == "")) {
			return false
		}
	}

	// Ensure we've consumed the entire string
	return pointer == length
}

func parsePattern(s string) []map[string]string {
	result := []map[string]string{}
	l := len(s)

	for i := 0; i < l; i++ {
		if s[i] == '*' {
			if i > 0 {
				result = append(result, map[string]string{
					string(s[i]): string(s[i-1]),
				})
			}
		} else if i+1 < l && (s[i+1] != '*') {
			result = append(result, map[string]string{
				"+": string(s[i]),
			})
		} else if i == l-1 {
			result = append(result, map[string]string{
				"+": string(s[i]),
			})
		}
	}

	return result
}

func main() {
	fmt.Println(isMatch("aab", "a*b"))
	//s = "aab", p = "c*a*b" is true

}
