package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var b func(start int) int
	b = func(s int) int {
		c := 0
		for s < len(nums) && nums[s] != 0 {
			c += 1
			s += 1
		}
		return c
	}

	consec := b(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c := b(i + 1)
			if consec < c {
				consec = c
			}
		}
	}

	return consec
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1}))
}
