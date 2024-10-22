package main

import "fmt"

func intToRoman(num int) string {
	m := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := ""

	for _, entry := range m {
		for entry.value <= num {
			roman += entry.symbol
			num -= entry.value
		}
	}

	return roman
}

func main() {
	fmt.Println(intToRoman(58))
}
