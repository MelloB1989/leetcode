package main

import "fmt"

func addBinary(a string, b string) string {
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	result := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		aa := 0
		bb := 0

		if i >= 0 && a[i] == '1' {
			aa = 1
		}
		if j >= 0 && b[j] == '1' {
			bb = 1
		}

		sum := aa + bb + carry
		if sum == 2 {
			carry = 1
			result = append([]byte{'0'}, result...)
		} else if sum == 3 {
			carry = 1
			result = append([]byte{'1'}, result...)
		} else {
			carry = 0
			result = append([]byte{byte(sum + '0')}, result...)
		}

		i--
		j--
	}

	return string(result)
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}
