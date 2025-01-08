package main

import "fmt"

// func findRepeatedDnaSequences(s string) []string {
// 	result := []string{}
// 	patternMap := map[string]bool{}
// 	check := func(c string, start int) bool {
// 		for i := start; i+10 <= len(s); i += 1 {
// 			// fmt.Println(c, s[i:i+10], c == s[i:i+10])
// 			if c == s[i:i+10] {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	for i := 0; i+10 < len(s); i++ {
// 		pattern := s[i : i+10]
// 		if check(pattern, i+1) && !patternMap[pattern] {
// 			result = append(result, pattern)
// 			patternMap[pattern] = true
// 		}
// 	}
// 	return result
// }

func findRepeatedDnaSequences(s string) []string {
	result := []string{}
	patternMap := map[string]int{}
	for i := 0; i+10 <= len(s); i++ {
		pattern := s[i : i+10]
		patternMap[pattern] += 1
	}
	for k, v := range patternMap {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
