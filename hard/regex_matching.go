package main

import "fmt"

func isMatch(s string, p string) bool {
	pointer := 0
	patterns := parsePattern(p)
	checks := 0
	// fmt.Print(patterns)
	length := len(s)
	r := true
	for i := 0; i < len(patterns) && pointer < length; i++ {
		for key, value := range patterns[i] {
			switch key {
			case "+":
				// Check for "." value
				if value == "." {
					pointer += 1
					// fmt.Println(value, " ", key, " ", r)
				} else if pointer < length && value == string(s[pointer]) { // Check if the current character matches
					pointer += 1
					// fmt.Println(value, " ", key, " ", r)
				} else {
					r = false // Fail if there's no match for "+"
					// fmt.Println(value, " ", key, " ", r)
				}
			case ".":
				pointer += 1 // Match any single character
			case "*":
				// Check for "." value
				if value == "." {
					r = true
					if len(patterns) > 1 {
						pointer = len(patterns) - 1
					} else {
						pointer = length
					}
				}
				// If the current character does not match the character before '*'
				if pointer < length && value != string(s[pointer]) {
					// Here, we can safely skip the '*'
					// We can just continue to the next pattern because '*' means zero occurrences
					r = true // Keep r true because `*` can represent zero occurrences
					// fmt.Println(value, " ", key, " ", r)
				} else {
					// If it matches, we match one or more occurrences
					for pointer < length && (value == string(s[pointer])) {
						pointer += 1
						// fmt.Println(value, " ", key, " ", r)
					}
				}
			}
			checks += 1
		}
	}
	// After processing all patterns, check if we have consumed all characters in s
	return r && pointer == length && checks == len(patterns)
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
	fmt.Println(isMatch("ab", ".*c"))
	//s = "aab", p = "c*a*b" is true

}
