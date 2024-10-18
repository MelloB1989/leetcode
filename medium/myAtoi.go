package main

import "fmt"

func myAtoi(s string) int {
	m := make(map[string]int)
	m["0"] = 0
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9

	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	var result, sign, i int
	sign = 1

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	for i < len(s) {
		if value, ok := m[string(s[i])]; ok {
			result = result*10 + value
			if sign*result > INT_MAX {
				return INT_MAX
			} else if sign*result < INT_MIN {
				return INT_MIN
			}
		} else {
			break
		}
		i++
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("    -042"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
}
