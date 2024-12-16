package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i]-i-1 >= k {
			return i + k
		}
	}
	return len(arr) + k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}
