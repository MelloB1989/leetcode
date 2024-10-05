package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLength := 0
	seen := make(map[byte]int)

	for end := 0; end < len(s); end++ {
		if idx, ok := seen[s[end]]; ok && idx >= start {
			start = idx + 1
		}
		seen[s[end]] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	fmt.Print(seen)
	return maxLength
}

func main() {
	s := "dvdf"
	fmt.Print(lengthOfLongestSubstring(s))
}
