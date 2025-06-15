package main

import "fmt"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	neighbourMap := map[int][]int{}
	var buildNeighbours func(currentEdge []int)
	buildNeighbours = func(currentEdge []int) {
		distTravelled := 0
		if currentEdge[2] <= distanceThreshold {
			neighbourMap[currentEdge[0]] = append(neighbourMap[currentEdge[0]], currentEdge[1])
			neighbourMap[currentEdge[1]] = append(neighbourMap[currentEdge[1]], currentEdge[0])
			distTravelled += currentEdge[2]
		}
		if distTravelled < currentEdge[2] {

		}
	}
	for _, edge := range edges {
		buildNeighbours(edge)
		// if edge[2] <= distanceThreshold {
		// 	neighbourMap[edge[0]] = append(neighbourMap[edge[0]], edge[1])
		// 	neighbourMap[edge[1]] = append(neighbourMap[edge[1]], edge[0])
		// }
	}
	min := n
	res := 0
	for k, v := range neighbourMap {
		if len(v) < min {
			res = k
			min = len(v)
		}
		fmt.Println(res, k, v)
		if len(v) == min && k > res {
			res = k
		}
	}
	return res

}

func main() {
	// fmt.Println(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
}
