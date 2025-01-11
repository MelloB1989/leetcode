package main

import "fmt"

// func zeroFilledSubarray(nums []int) int64 {
// 	preSum := map[int]int{}
// 	count := 0
// 	preSum[0] = 1
// 	sum := 0

// 	for i := 0; i < len(nums); i++ {
// 		sum += abs(nums[i])

// 		if f, exists := preSum[sum]; exists {
// 			count += f
// 		}

// 		preSum[sum] += 1
// 	}

// 	return int64(count)
// }

//	func abs(x int) int {
//		if x < 0 {
//			return -x
//		}
//		return x
//	}
func zeroFilledSubarray(nums []int) int64 {
	var zc int64
	var total int64

	for _, v := range nums {
		if v == 0 {
			total = total + zc + 1
			zc += 1
		} else {
			zc = 0
		}
	}
	return total
}

func main() {
	fmt.Println(zeroFilledSubarray([]int{0, -9, 6, -5, 0, 0, 8, 0, 0, 3, -3}))
}
