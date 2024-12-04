package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	k := 5

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
