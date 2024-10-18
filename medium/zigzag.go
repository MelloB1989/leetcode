package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	m := make(map[int][]rune)
	r := 1
	zig := 1
	sf := ""
	for i := 0; i <= len(s)-1; i++ {
		m[r] = append(m[r], rune(s[i]))
		if zig == 1 {
			r++
		} else {
			r--
		}
		if r == numRows || r == 1 {
			if zig == 0 {
				zig = 1
			} else {
				zig = 0
			}
		}
	}
	for i := 1; i <= numRows; i++ {
		sf = sf + string(m[i])
	}
	return sf
}

func main() {
	s := "AB"
	numRows := 1
	fmt.Print(convert(s, numRows))
}
