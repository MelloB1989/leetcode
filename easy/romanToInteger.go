package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[string]int)
	result := 0
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	for i := 0; i < len(s); i++ {
		if i+1 <= len(s)-1 && m[string(s[i])] < m[string(s[i+1])] {
			result -= m[string(s[i])]
		} else {
			result += m[string(s[i])]
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
