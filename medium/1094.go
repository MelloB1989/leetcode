package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	stops := make([]int, 1001)
	for _, trip := range trips {
		n, from, to := trip[0], trip[1], trip[2]
		stops[from] += n
		stops[to] -= n
	}
	for currentPassengers, i := 0, 0; i < len(stops); i++ {
		currentPassengers += stops[i]
		if currentPassengers > capacity {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(carPooling([][]int{
		{2, 1, 5},
		{3, 3, 7},
	}, 4)) // true for 5
}
