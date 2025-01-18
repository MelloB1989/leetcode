package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	result = append(result, intervals[0])
	rp := 0
	for i := 1; i < len(intervals); i++ {
		prev := result[rp]
		curr := intervals[i]

		if prev[1] >= curr[0] {
			result[rp][0] = min(prev[0], curr[0])
			result[rp][1] = max(prev[1], curr[1])
			// result[rp][1] = curr[1]
			// if prev[0] >= curr[0] {
			// 	result[rp][0] = curr[0]
			// }
		} else {
			result = append(result, intervals[i])
			rp += 1
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
}
