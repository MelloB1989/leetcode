package main

import "fmt"

func isMatch(s string, p string) bool {
	pointer := 0
	patterns := parsePattern(p)
	fmt.Print(patterns)
	checks := 0
	length := len(s)
	r := true

	for i := 0; i < len(patterns) && pointer < length; i++ {
		for key, value := range patterns[i] {
			switch key {
			case "+":
				// Ensure at least one match for "+"
				if pointer < length && (value == string(s[pointer]) || value == ".") {
					// Move pointer forward while there are matches
					pointer += 1
					fmt.Println(value, " ", key, " ", r)
				} else {
					r = false // Fail if there's no match for "+"
					fmt.Println(value, " ", key, " ", r)
				}
			case ".":
				pointer += 1 // Match any single character
			case "*":
				// If the current character does not match the character before '*'
				if pointer < length && value != string(s[pointer]) {
					// `*` can match zero occurrences
					r = true // Keep r true because `*` can represent zero occurrences
					fmt.Println(value, " ", key, " ", r)
				} else {
					// If it matches, we match one or more occurrences
					for pointer < length && (value == string(s[pointer]) || value == ".") {
						pointer += 1
						fmt.Println(value, " ", key, " ", r)
					}
				}
			}
			checks += 1
		}
	}

	// After processing all patterns, check if we have consumed all characters in s
	fmt.Println("pointer: ", pointer)
	return r && pointer == length && checks == len(patterns) // Ensure that all characters in s are matched
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
	fmt.Println(isMatch("aaa", "a*a"))
	//s = "aab", p = "c*a*b" is true

}
