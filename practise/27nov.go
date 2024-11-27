package main

import "fmt"

// Sliding window
func main() {
	k := 4
	arr := []int{8, -7, 2, 5, 4, -3, 3, -2, -6}
	max := 0
	// n := 0
	// for i := 0; i < len(arr)-k; i++ {
	// 	sum := 0
	// 	sum += arr[i]
	// 	for j := 1; j < k; j++ {
	// 		sum += arr[i+j]
	// 		n++
	// 	}
	// 	if sum > max {
	// 		max = sum
	// 	}
	// 	n++
	// }
	sum := 0
	n := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
		n++
	}
	max = sum
	for i := k; i < len(arr); i++ {
		sum += arr[i]
		sum -= arr[i-k]
		if sum > max {
			max = sum
		}
		n++
	}

	fmt.Println(max, n)
}

//min subarray, max no of vowels, max no of odd no. in subarrray
