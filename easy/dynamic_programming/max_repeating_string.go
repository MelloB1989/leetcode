package main

import (
	"fmt"
)

func maxRepeating(sequence string, word string) int {
	length := len(sequence)
	lenght_word := len(word)
	if length <= 1 && word == sequence {
		return 1
	}
	max := 0
	i := 0
	for i < length {
		if i+lenght_word < length && sequence[i:i+lenght_word] == word {
			max += 1
			// i += lenght_word - 1
		}
		i++
	}
	return max
}

func main() {
	fmt.Println(maxRepeating("ababc", "ab"))
}
