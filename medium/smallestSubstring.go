package main

import "fmt"

// func smallest(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}

// 	charSet := make(map[rune]bool)
// 	for _, c := range s {
// 		charSet[c] = true
// 	}

// 	minLen := len(s)

// 	for i := 0; i < len(s); i++ {
// 		currentChars := make(map[rune]bool)
// 		for j := i; j < len(s); j++ {
// 			currentChars[rune(s[j])] = true
// 			if len(currentChars) == len(charSet) {
// 				if j-i+1 < minLen {
// 					minLen = j - i + 1
// 				}
// 				break
// 			}
// 		}
// 	}

//		return minLen
//	}
func smallest(s string) int {
	uniqueChars := make(map[rune]bool)
	for _, c := range s {
		uniqueChars[c] = true
	}
	uniqueCount := len(uniqueChars)

	charFreq := make(map[rune]int)
	start := 0
	minLen := len(s) + 1
	count := 0

	for end := 0; end < len(s); end++ {
		curr := rune(s[end])
		charFreq[curr]++
		if charFreq[curr] == 1 {
			count++
		}

		if count == uniqueCount {
			for count == uniqueCount {
				if end-start+1 < minLen {
					minLen = end - start + 1
				}

				startChar := rune(s[start])
				charFreq[startChar]--
				if charFreq[startChar] == 0 {
					count--
				}
				start++
			}
		}
	}

	if minLen > len(s) {
		return 0
	}
	return minLen
}

func main() {
	fmt.Println(smallest("abcaabddcdbbbbc"))
}
