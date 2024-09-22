package main

import "fmt"

func isPalindrome(x int) bool {
	var rev = 0
	var original = x
	for x > 0 {
		rev = 10*rev + x%10
		x = x / 10
	}
	if rev == original {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
}
