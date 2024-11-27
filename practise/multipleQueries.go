package main

import "fmt"

// Multiple Queries Sum
func main() {
	arr := []int{7, 1, 3, 4, 5, 9, 2, 8}
	queries := [][]int{
		{0, 5},
		{2, 6},
		{2, 4},
		{3, 7},
		{1, 7},
	}
	prefixSum := make([]int, len(arr)+1)
	result := make([]int, len(queries))

	for i := 1; i <= len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i-1]
	}

	for i, query := range queries {
		l, r := query[0], query[1]
		if l == 0 {
			result[i] = prefixSum[r]
		} else {
			result[i] = prefixSum[r] - prefixSum[l-1]
		}
	}

	fmt.Println(result)
}
