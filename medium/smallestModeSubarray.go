package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 1, 4, 2}
	minLen := len(arr)
	mode := findMode(arr)
	for i := 0; i < len(arr); i++ {
		freqMap := make(map[int]int)
		for j := i; j < len(arr); j++ {
			freqMap[arr[j]] += 1
			sm := 0
			for _, count := range freqMap {
				if count > sm {
					sm = count
				}
			}
			subArrayLen := j - i + 1
			if subArrayLen < minLen && sm == mode {
				minLen = subArrayLen
			}
		}
	}
	fmt.Println(minLen)
}

func findMode(arr []int) int {
	mode := arr[0]
	fm := make(map[int]int)
	for _, num := range arr {
		fm[num]++
	}
	for _, count := range fm {
		if count > mode {
			mode = count
		}
	}
	return mode
}
