package main

import "fmt"

func rotate(nums []int, k int) {
	var r func(n []int) []int
	r = func(n []int) []int {
		last := n[len(n)-1]
		for i := len(n) - 1; i > 0; i-- {
			n[i] = n[i-1]
		}
		n[0] = last
		return n
	}

	for i := 0; i < k; i++ {
		nums = r(nums)
	}

	fmt.Println(nums)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3) //7,1,2,3,4,5,6
}
