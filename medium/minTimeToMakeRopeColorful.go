package main

import (
	"fmt"
)

// 875,1642
func minCost(colors string, neededTime []int) int {
	var removeColors func(times []int) int
	removeColors = func(times []int) int {
		if len(times) == 0 {
			return 0
		}
		minTime := 0
		maxTime := 0
		for _, time := range times {
			minTime += time
			if time > maxTime {
				maxTime = time
			}
		}
		minTime -= maxTime
		return minTime
	}
	totalMinTime := 0
	n := len(colors)
	i := 0

	for i < n {
		color := colors[i]
		j := i + 1
		for j < n && colors[j] == color {
			j++
		}
		// Now, colors[i:j] are all the same
		if j-i > 1 {
			// Segment from i to j-1
			fmt.Println("Not colorful from", i, "to", j-1)
			times := neededTime[i:j]
			fmt.Println("Times:", times)
			minTime := removeColors(times)
			fmt.Println("MinTime:", minTime)
			totalMinTime += minTime
		}
		i = j
	}
	return totalMinTime
}

func main() {
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1, 1}))
}
