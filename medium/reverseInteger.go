package main

import "fmt"

func reverse(x int) int {
	const maxInt = 1<<31 - 1 // 2147483647
	const minInt = -1 << 31  // -2147483648

	result := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if result > maxInt/10 || (result == maxInt/10 && pop > 7) {
			return 0
		}
		if result < minInt/10 || (result == minInt/10 && pop < -8) {
			return 0
		}

		result = result*10 + pop
	}

	return result
}

func main() {
	fmt.Println(reverse(153423))
	// reverse(-123)
}
