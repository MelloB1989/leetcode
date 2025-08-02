package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sumSet := []int{}
	currentSum := 0
	var b func(e int)

	sort.Ints(candidates)

	b = func(e int) {
		if currentSum == target {
			s := make([]int, len(sumSet))
			copy(s, sumSet)
			result = append(result, s)
		}
		for i := e; i < len(candidates); i++ {
			if currentSum+candidates[i] <= target {
				if i > e && candidates[i] == candidates[i-1] {
					continue
				}
				currentSum += candidates[i]
				sumSet = append(sumSet, candidates[i])
				b(i + 1)
				currentSum -= candidates[i]
				sumSet = sumSet[:len(sumSet)-1]
			}
		}
	}

	b(0)
	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
