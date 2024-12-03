package main

import "fmt"

func main() {
	arr := []int{5, -2, -1, 3, 2, 5, -4, 8}
	k := 3
	fmt.Println(subarrayDivisibleByK(arr, k))
}

func subarrayDivisibleByK(arr []int, k int) int {
	preSum := make([]int, len(arr)+1)
	count := 0
	remainder := make([]int, k)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + arr[i-1]
		remainder[(preSum[i]%k+k)%k]++
	}
	remainder[0]++
	// for i := 0; i < len(preSum); i++ {
	// 	remainder[(preSum[i]%k+k)%k]++
	// }

	for i := 0; i < k; i++ {
		count += (remainder[i] * (remainder[i] - 1)) / 2
	}
	fmt.Println(remainder)

	return count
}
