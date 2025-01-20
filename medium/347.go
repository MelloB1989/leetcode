package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	type pair struct {
		key   int
		value int
	}
	pairs := make([]pair, 0, len(freqMap))

	for key, value := range freqMap {
		pairs = append(pairs, pair{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	result := make([]int, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		result = append(result, pairs[i].key)
	}

	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 2))
	// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
