package main

import "fmt"

func getLargestOutlier(nums []int) int {
	sum := 0
	outlier := nums[0]
	frequency := make(map[int]int)
	for _, num := range nums {
		sum += num
		frequency[num*2]++
	}

	for _, num := range nums {
		t := sum - num
		if frequency[t] >= 2 || (frequency[t] == 1 && t != num*2) {
			if num > outlier {
				outlier = num
			}
		}
	}
	return outlier
}

func main() {
	nums := []int{2, 3, 5, 10}
	result := getLargestOutlier(nums)
	fmt.Println(result)
}
