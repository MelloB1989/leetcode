package main

import "fmt"

func reverseString(s string) string {
	if len(s) == 0 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}

func checkPalindrom(s string) bool {
	return reverseString(s) == s
}

func longestPalindrome(s string) string {
	length := len(s)
	longest := string(s[0])
	if length == 0 {
		return ""
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			sub := s[i : j+1]
			if checkPalindrom(sub) && len(sub) > len(longest) {
				longest = sub
			}
		}
	}
	return longest
}

func main() {
	s := "bb"
	fmt.Print(longestPalindrome(s))
}
