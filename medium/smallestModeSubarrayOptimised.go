package main

import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 2, 2, 1, 5, 1, 5, 6, 7, 8}
// 	freqMap := make(map[int]int)
// 	fo := make(map[int]int)
// 	lo := make(map[int]int)
// 	for i, num := range arr {
// 		if _, exists := freqMap[num]; !exists {
// 			fo[num] = i
// 		}
// 		freqMap[num]++
// 		lo[num] = i
// 	}
// 	max := 0
// 	for _, freq := range freqMap {
// 		if freq > max {
// 			max = freq
// 		}
// 	}

//		minLen := len(arr)
//		for num, freq := range freqMap {
//			if freq == max {
//				l := lo[num] - fo[num] + 1
//				if l < minLen {
//					minLen = l
//				}
//			}
//		}
//		fmt.Print(minLen)
//	}
func main() {
	arr := []int{1, 2, 3, 2, 2, 1, 5, 1, 5, 6, 7, 8}
	freqMap := make(map[int]int)
	oc := make(map[int][]int)
	for i, num := range arr {
		oc[num] = append(oc[num], i)
		freqMap[num]++
	}
	max := 0
	for _, freq := range freqMap {
		if freq > max {
			max = freq
		}
	}

	minLen := len(arr)
	fmt.Println(oc)
	for num, freq := range freqMap {
		if freq == max {
			l := oc[num][(len(oc[num])-1)] - oc[num][0] + 1
			if l < minLen {
				minLen = l
			}
		}
	}
	fmt.Print(minLen)
}
