package main

import "fmt"

// func rearrangeArray(nums []int) []int {
// 	resultArray := make([]int, len(nums))

// 	getNext := func(neg bool) int {
// 		for i := 0; i < len(nums); i++ {
// 			if nums[i] == 0 {
// 				continue
// 			}
// 			if neg {
// 				if nums[i] < 0 {
// 					p := nums[i]
// 					nums[i] = 0
// 					return p
// 				}
// 			} else {
// 				if nums[i] > 0 {
// 					p := nums[i]
// 					nums[i] = 0
// 					return p
// 				}
// 			}
// 		}
// 		return 0 //should never reach here
// 	}

// 	resultArray[0] = getNext(false)

// 	for i := 1; i < len(resultArray); i++ {
// 		neg := true
// 		if resultArray[i-1] < 0 {
// 			neg = false
// 		}

// 		resultArray[i] = getNext(neg)
// 	}

//		return resultArray
//	}
func rearrangeArray(nums []int) []int {
	resultArray := make([]int, len(nums))
	pos := 0
	neg := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			resultArray[neg] = nums[i]
			neg += 2
		} else {
			resultArray[pos] = nums[i]
			pos += 2
		}
	}

	return resultArray
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
}
