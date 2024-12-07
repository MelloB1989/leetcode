package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	rot := make([]int, n)

	for i := 0; i < n; i++ {
		rot[(i+k)%n] = nums[i]
	}

	for i := 0; i < n; i++ {
		nums[i] = rot[i]
	}

	fmt.Println(nums)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3) //7,1,2,3,4,5,6
}
