package main

import "fmt"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	helpNeeded := map[int]int{}
	for i := 0; i < len(heights)-1; i++ {
		if heights[i] < heights[i+1] {
			helpNeeded[i] = (heights[i+1] - heights[i])
		}
	}

	fmt.Println(helpNeeded)

	for place, help := range helpNeeded {

	}

	return 0
}

func main() {
	fmt.Println(furthestBuilding([]int{4, 8, 2, 4, 6, 8}, 5, 1))
	// fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
}
