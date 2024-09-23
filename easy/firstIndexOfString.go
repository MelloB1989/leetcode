package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}
