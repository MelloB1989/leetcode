package main

import "fmt"

func mostBooked(n int, meetings [][]int) int {
	meetingsHeld := make([]int, n)
	for
}

func main() {
	fmt.Println(mostBooked(2, [][]int{
		{0, 10},
		{1, 5},
		{2, 7},
		{3, 4},
	}))
}
