package main

import "fmt"

func main() {
	arr := []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}
	k := 2

	fmt.Println(maxOperations(arr, k))
}

func maxOperations(arr []int, k int) int {
	done := map[int]int{}
	count := 0
	for _, num := range arr {
		fmt.Println(done)
		complement := k - num
		fmt.Println(complement)
		if found := done[complement]; found > 0 {
			count++
			done[complement]--
		} else {
			done[num]++
		}
	}
	return count
}
