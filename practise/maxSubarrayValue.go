package main

import "fmt"

func main() {
	k := 4
	arr := []int{9, -8, 6, 9, 2, -5, 2, 7, 8, 4}
	sum := 0
	max := 0
	count := make(map[int]int)
	subCount := 0
	m := arr[0]
	fm := -1
	for i := 0; i < k; i++ {
		sum += arr[i]
		count[arr[i]] += 1
		if count[arr[i]] == 1 {
			subCount++
		}
		if m < arr[i] {
			m = arr[i]
		}
	}
	if subCount == k {
		max = sum
		fm = m
	}
	m = arr[0]
	for i := k; i < len(arr); i++ {
		sum -= arr[i-k]
		count[arr[i-k]]--
		if count[arr[i-k]] == 0 {
			subCount--
		}
		sum += arr[i]
		count[arr[i]] += 1
		if count[arr[i]] == 1 {
			subCount++
		}
		if subCount == k && sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
