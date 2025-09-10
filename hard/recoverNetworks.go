package main

import (
	"fmt"
	"sort"
)

func recoverNetwork(networkNodes int32, networkFrom []int32, networkTo []int32, company int32) []int32 {
	adj := make(map[int32][]int32)
	for i := 0; i < len(networkFrom); i++ {
		from := networkFrom[i]
		to := networkTo[i]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	visited := map[int32]bool{}
	visited[company] = true
	queue := []int32{company}
	result := []int32{}

	for len(queue) > 0 {
		current_node := queue[0]
		queue = queue[1:]
		current_level := []int32{}
		for _, neighbour := range adj[current_node] {
			if !visited[neighbour] {
				visited[neighbour] = true
				current_level = append(current_level, neighbour)
			}
		}
		sort.Slice(current_level, func(i int, j int) bool {
			return current_level[i] < current_level[j]
		})
		queue = append(queue, current_level...)
		result = append(result, current_level...)
	}
	return result
}

func main() {
	fmt.Println(recoverNetwork(5, []int32{1, 1, 2, 3, 1}, []int32{2, 3, 4, 5, 5}, 1))
}
