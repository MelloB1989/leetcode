package main

import "fmt"

// max and second max
// func main() {
// 	arr := []int{9, 7, 8}
// 	m := arr[0]
// 	s := arr[1]
// 	if arr[1] > m {
// 		m = arr[1]
// 		s = arr[0]
// 	}
// 	for i := 2; i < len(arr); i++ {
// 		if arr[i] > m {
// 			s = m
// 			m = arr[i]
// 		} else if arr[i] > s {
// 			s = arr[i]
// 		}
// 	}
// 	fmt.Print(m, s)
// }

// Print all sub arrays
// func main() {
// 	// arr := []int{-8, 7, 2, 5, -6, -1, 7, 2, -6, 4}
// 	arr := []int{-8, -9, -7}
// 	max := 0
// 	for i := 0; i < len(arr); i++ {
// 		for j := i; j < len(arr); j++ {
// 			sum := 0
// 			for k := i; k <= j; k++ {
// 				// fmt.Print(arr[k], " ")
// 				sum += arr[k]
// 			}
// 			if max == 0 {
// 				max = sum
// 			} else if sum > max {
// 				max = sum
// 			}
// 			// fmt.Println("")
// 		}
// 		// fmt.Println("")
// 	}
// 	fmt.Print(max)
// }

// Max sub array
// func main() {
// 	// arr := []int{-8, 7, 2, 5, -6, -1, 7, 2, -6, 4}
// 	arr := []int{-8, -9, -7}
// 	max := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum := 0
// 		for j := i; j < len(arr); j++ {
// 			sum = sum + arr[j]
// 			if max == 0 {
// 				max = sum
// 			} else if sum > max {
// 				max = sum
// 			}
// 		}
// 	}
// 	fmt.Print(max)
// }

// Kadanes algorithm
//
//	func main() {
//		arr := []int{-8, -9, -7}
//		localMax := arr[0]
//		sum := arr[0]
//		for i := 0; i < len(arr); i++ {
//			sum += arr[i]
//			if sum > localMax {
//				localMax = sum
//			}
//			if sum < 0 {
//				sum = 0
//			}
//		}
//		fmt.Print(localMax)
//	}
//
// Length of longest subarray with odd sum
func main() {
	arr := []int{6, 7, 2, 5, 1, 8, 6, 9, 2, 2}
	maxLength := 0
	currentSum := 0
	// oddIndex := 0
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		if arr[i]%2 != 0 && n > i+1 {
			// oddIndex = i
			n = i + 1
		}
	}
	maxLength = len(arr) - n
	fmt.Print(maxLength)
}

// o+e
// func main() {
// 	// arr := []int{-8, 7, 2, 5, -6, -1, 7, 2, -6, 4}
// 	arr := []int{6, 7, 2, 5, 1, 8, 6, 9, 2, 2}
// 	max := 0
// 	l := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum := 0
// 		for j := i; j < len(arr); j++ {
// 			sum = sum + arr[j]
// 			if max == 0 {
// 				max = sum
// 			} else if sum > max && sum%2 != 0 {
// 				max = sum
// 				l = (j - i)
// 			}
// 		}
// 	}
// 	fmt.Print(max, l)
// }
