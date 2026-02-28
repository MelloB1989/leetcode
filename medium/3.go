package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	unique := make(map[rune]int)
	maxLen := 0
	left := 0
	for right := 0; right < n; right++ {
		unique[rune(s[right])]++
		for unique[rune(s[right])] > 1 {
			unique[rune(s[left])]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
