package main

import "fmt"

func addBinary(a string, b string) string {
	carry := 0
	result := []byte("0000")
	n := 4 - len(a)
	for i := 0; i < n; i++ {
		a += "0"
	}

	m := 4 - len(b)
	for i := 0; i < m; i++ {
		b += "0"
	}
	for i := 3; i >= 0; i-- {
		aa := 0
		bb := 0
		if string(a[i]) == "1" {
			aa = 1
		}
		if string(b[i]) == "1" {
			bb = 1
		}
		sum := aa + bb + carry
		if sum == 2 {
			carry = 1
			result[i] = '0'
		} else if sum == 3 {
			carry = 1
			result[i] = '1'
		} else {
			result[i] = byte(sum + '0')
		}
	}
	for i := 0; i < len(result)-2; i++ {
		if result[i] == '0' {
			result = result[1:]
		}
	}
	return string(result)
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}
