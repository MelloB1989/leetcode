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
	if length <= 1 {
		return s
	}
	maxLen := 1
	maxStr := s[0:1]
	for i := 0; i < length; i++ {
		for j := i + maxLen; j < length; j++ {
			fmt.Print(j, " ", i, " ", maxLen, " ", maxStr, " ", s[i:j], "\n")
			if checkPalindrom(s[i:j]) && j-i > maxLen {
				maxLen = j - 1
				maxStr = s[i:j]
			}
		}
	}
	return maxStr
}

func main() {
	s := "babad"
	fmt.Print(longestPalindrome(s))
}
