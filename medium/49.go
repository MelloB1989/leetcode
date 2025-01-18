package main

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	createSig := func(s string) string {
		count := make([]int, 26)

		for _, char := range s {
			count[char-'a'] += 1
		}

		var builder strings.Builder
		for _, c := range count {
			builder.WriteString(strconv.Itoa(c))
			builder.WriteByte('#')
		}

		return builder.String()
	}
	result := [][]string{}
	sig := map[string][]string{}
	for _, word := range strs {
		s := createSig(word)
		sig[s] = append(sig[s], word)
	}

	for _, v := range sig {
		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
}
