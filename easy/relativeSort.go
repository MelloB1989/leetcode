package main

import "fmt"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	freqMap := make(map[int]int)
	result := make([]int, 0, len(arr1))

	for _, num := range arr1 {
		freqMap[num]++
	}

	for _, num := range arr2 {
		if freqMap[num] > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	for num, freq := range freqMap {
		if freq > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	return sortArray(result)
}

func sortArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	result := relativeSortArray(arr1, arr2)
	fmt.Println(result)
}
