package main

import "fmt"

func minCostConnectPoints(points [][]int) int {
	const maxInt = 1<<31 - 1
	cost := 0
	np := len(points)
	visited := make([]bool, np)
	dist := make([]int, np)
	for i := 0; i < np; i++ {
		dist[i] = maxInt
	}
	dist[0] = 0
	getMinDistVertex := func() int {
		minDist := maxInt
		minVertex := -1

		for i := 0; i < len(dist); i++ {
			if !visited[i] && dist[i] < minDist {
				minDist = dist[i]
				minVertex = i
			}
		}

		return minVertex
	}
	for i := 0; i < np; i++ {
		u := getMinDistVertex()
		visited[u] = true
		cost += dist[u]

		for v := 0; v < np; v++ {
			d := abs(points[u][0]-points[v][0]) + abs(points[u][1]-points[v][1])
			dist[v] = min(dist[v], d)
		}
	}
	return cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println(minCostConnectPoints(points))
}
