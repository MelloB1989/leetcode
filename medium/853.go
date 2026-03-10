package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	fleets := 0
	idx := make([]int, len(position))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return position[idx[i]] > position[idx[j]]
	})
	var lastTime float64
	for i := range idx {
		var time float64
		time = float64(target-position[idx[i]]) / float64(speed[idx[i]])
		if i == 0 {
			lastTime = time
			fleets++
			continue
		}
		if time > lastTime {
			fleets++
			lastTime = time
		}
	}
	return fleets
}

func main() {
	// fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{6, 8}, []int{3, 2}))
}
