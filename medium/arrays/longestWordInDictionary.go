package main

import "fmt"

func findLongestWord(s string, dictionary []string) string {
	freqMap := make(map[string]int)
	for _, c := range s {
		freqMap[string(c)]++
	}
	for i := 0; i < len(dictionary); i++ {

	}
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
