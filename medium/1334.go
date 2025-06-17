package main

import "fmt"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	const INT_MAX = 1<<31 - 1
	visited := make([]int, n)
	dist := make([]int, n)
	queue := make([]int, n)
	for i, _ := range dist {
		dist[i] = INT_MAX
	}
	graph := [][]int{}

	var bfs func()
	bfs = func() {

	}
}

func main() {
	// fmt.Println(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
}
