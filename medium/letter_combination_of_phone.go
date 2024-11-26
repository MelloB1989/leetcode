package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	letterMap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{""}
	}
	var result []string
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			result = append(result, path)
			return
		}
		letters := letterMap[string(digits[index])]
		for _, letter := range letters {
			fmt.Println(path, letter)
			backtrack(index+1, path+letter)
		}
	}
	backtrack(0, "")
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
