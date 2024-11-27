package main

import "fmt"

// func main() {
// 	k := 4
// 	arr := []int{9, -8, 6, 9, 2, -5, 2, 7, 8, 4}
// 	subStart := 0
// 	sum := 0
// 	unique := true
// 	max := 0

// 	em := make(map[int]bool)
// 	for i := 0; i < k; i++ {
// 		sum += arr[i]
// 		if em[arr[i]] {
// 			unique = false
// 		}
// 		em[arr[i]] = true
// 	}
// 	max = sum
// 	for i := k; i < len(arr); i++ {
// 		sum -= arr[i-k]
// 		if em[arr[i-k]] {
// 			delete(em, arr[i-k])
// 		}
// 		sum += arr[i]
// 		if em[arr[i]] {
// 			unique = false
// 		} else {
// 			unique = true
// 		}
// 		if sum > max && unique {
// 			max = sum
// 			subStart = i - k + 1
// 		}
// 	}
// 	fmt.Println(max, subStart, unique)
// }

//Using count

func main() {
	k := 4
	arr := []int{9, -8, 6, 9, 2, -5, 2, 7, 8, 4}
	sum := 0
	max := 0
	count := make(map[int]int)
	subCount := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
		count[arr[i]] += 1
		if count[arr[i]] == 1 {
			subCount++
		}
	}
	if subCount == k {
		max = sum
	}
	for i := k; i < len(arr); i++ {
		sum -= arr[i-k]
		// if count[arr[i-k]] >= 1 {
		// 	delete(count, arr[i-k])
		// }
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
