package main

import "fmt"

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if sum == 10 {
			result[i+1] = 0
			carry = 1
		} else {
			result[i+1] = sum
			carry = 0
		}
	}
	if carry == 1 {
		result[0] = 1
		return result
	}

	return result[1:]
}

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}
