package main

import "fmt"

func isAnagram(w1, w2 string) bool {
	m := map[rune]int{}
	for _, v := range w1 {
		m[v] += 1
	}
	for _, v := range w2 {
		m[v] -= 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagram := [][]string{}
	for i := 0; i < len(words); i++ {
		currentList := []string{}
		currentList = append(currentList, words[i])
		for j := i + 1; j < len(words); j++ {
			if isAnagram(words[i], words[j]) {
				currentList = append(currentList, words[j])
			}
		}
		anagram = append(anagram, currentList)
	}

	fmt.Println(anagram)
}
