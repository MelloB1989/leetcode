package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 12, 18, 25, 32}
	x := 19
	minMax := -1

	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2

	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[mid])
		if arr[mid] == x {
			minMax = -1
			break
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			minMax = arr[mid]
			right = mid - 1
		}
		mid = (left + right) / 2
	}

	fmt.Println(minMax)
}

// func main() {
// 	arr := []int{1, 3, 4, 12, 18, 25, 32}
// 	x := 3
// 	minMax := -1

// 	left := 0
// 	right := len(arr) - 1
// 	mid := (left + right) / 2

// 	for i := 0; i < len(arr); i++ {
// 		if arr[mid] < x {
// 			left = mid + 1
// 		} else if arr[mid] > x {
// 			if minMax == -1 || minMax < arr[mid] {
// 				minMax = arr[mid]
// 			}
// 			// right = mid - 1
// 		}
// 		mid = (left + right) / 2
// 	}

// 	fmt.Println(minMax)
// }
