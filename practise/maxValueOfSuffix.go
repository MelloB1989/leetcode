package main

import "fmt"

// func main() {
// 	arr := []int{7, -3, 2, -5, 4, 3}
// 	for i := 0; i < len(arr); i++ {
// 		//Suffix array
// 		max := 0
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[j] > max {
// 				max = arr[j]
// 			}
// 		}
// 		if i == len(arr)-1 {
// 			max = -1
// 		}
// 		fmt.Print(max, " ")
// 	}
// }
//More optimal

func main() {
	arr := []int{7, -3, 2, -5, 4, 3}
	max := arr[len(arr)-1]
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max && i != 0 {
			max = arr[i]
		}
		if i == len(arr)-1 {
			fmt.Print(-1, " ")
		} else {
			fmt.Print(max, " ")
		}
	}
}
