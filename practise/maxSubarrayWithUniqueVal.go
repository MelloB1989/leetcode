package main

import "fmt"

func main() {
	k := 4
	arr := []int{9, -8, 6, 9, 2, -5, 2, 7, 8, 4}
	subStart := 0
	sum := 0
	unique := true
	max := 0

	em := make(map[int]bool)
	for i := 0; i < k; i++ {
		sum += arr[i]
		if em[arr[i]] {
			unique = false
		}
		em[arr[i]] = true
	}
	max = sum
	for i := k; i < len(arr); i++ {
		sum -= arr[i-k]
		if em[arr[i-k]] {
			delete(em, arr[i-k])
		}
		sum += arr[i]
		if em[arr[i]] {
			unique = false
		} else {
			unique = true
		}
		if sum > max && unique {
			max = sum
			subStart = i - k + 1
		}
	}
	fmt.Println(max, subStart, unique)
}
