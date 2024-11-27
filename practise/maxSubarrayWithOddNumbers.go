package main

import "fmt"

func main() {
	max := -1
	k := 4
	arr := []int{9, -8, 6, 9, 2, -5, 3, 7, 9, 11}
	odds := 0
	sum := 0

	for i := 0; i < k; i++ {
		sum += arr[i]
		if arr[i]%2 != 0 {
			odds += 1
		}
	}
	if max < sum && odds == k {
		max = sum
	}
	for i := k; i < len(arr); i++ {
		sum += arr[i]
		sum -= arr[i-k]
		if arr[i-k]%2 != 0 {
			odds -= 1
		}
		if arr[i]%2 != 0 {
			odds += 1
		}
		fmt.Println(sum, odds)
		if sum > max && odds == k {
			max = sum
		}
	}
	fmt.Println(max)
}
