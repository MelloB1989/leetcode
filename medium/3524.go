package main

import "fmt"

func resultArray(nums []int, k int) []int64 {
	// m := map[int]int{}
	result := make([]int64, k)
	for i, _ := range nums {
		pr := 1
		for j := i; j < len(nums); j++ {
			pr = (nums[j+1] * pr) % k
			result[pr] += 1
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	fmt.Println(resultArray(nums, k))
}
