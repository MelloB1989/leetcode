package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool {
		return r[i] > r[j]
	})
	return string(r)
}

func groupAnagrams(words []string) [][]string {
	m := map[string][]string{}
	for _, word := range words {
		sortedWord := sortString(word)
		m[sortedWord] = append(m[sortedWord], word)
	}

	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	// words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	words := []string{"", ""}
	anagrams := groupAnagrams(words)
	fmt.Println(anagrams)
}

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
