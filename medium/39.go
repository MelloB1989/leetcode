package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	sumSet := []int{}
	currentSum := 0

	var b func(e int)
	b = func(e int) {
		if currentSum == target {
			s := make([]int, len(sumSet))
			copy(s, sumSet)
			result = append(result, s)
		}
		for i := e; i < len(candidates); i++ {
			if currentSum+candidates[i] <= target {
				currentSum += candidates[i]
				sumSet = append(sumSet, candidates[i])
				b(i)
				currentSum -= sumSet[len(sumSet)-1]
				sumSet = sumSet[:len(sumSet)-1]
			}
		}
	}
	b(0)
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	// fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
