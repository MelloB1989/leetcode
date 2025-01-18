package main

import "fmt"

func findMaxLength(nums []int) int {
	zc := 0
	oc := 0
	max := 0
	m := map[int]int{}

	m[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zc += 1
		} else {
			oc += 1
		}

		if v, exists := m[zc-oc]; exists {
			if i-v > max {
				max = i - v
			}
		} else {
			m[zc-oc] = i
		}
	}

	return max
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}
