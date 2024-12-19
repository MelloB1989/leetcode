package main

import (
	"fmt"
)

func minGroups(intervals [][]int) int {
	n := len(intervals)
	unprocessed := make([]int, n)
	for i := 0; i < n; i++ {
		unprocessed[i] = i
	}

	count := 0

	for len(unprocessed) > 0 {
		firstKey := unprocessed[0]
		s := intervals[firstKey][0]
		e := intervals[firstKey][1]
		unprocessed = unprocessed[1:]

		toDelete := []int{}
		for i, k := range unprocessed {
			st := intervals[k][0]
			end := intervals[k][1]
			if end < s || st > e {
				toDelete = append(toDelete, i)
				if st < s {
					s = st
				}
				if end > e {
					e = end
				}
			}
		}

		for i := len(toDelete) - 1; i >= 0; i-- {
			idx := toDelete[i]
			unprocessed = append(unprocessed[:idx], unprocessed[idx+1:]...)
		}

		count++
	}

	return count
}

func main() {
	fmt.Println(minGroups([][]int{
		{441459, 446342}, {801308, 840640}, {871890, 963447}, {228525, 336985}, {807945, 946787}, {479815, 507766}, {693292, 944029}, {751962, 821744},
	}))
	// fmt.Println(minGroups([][]int{
	// 	{1, 3}, {5, 6}, {8, 10}, {11, 13},
	// }))
	// fmt.Println(minGroups([][]int{
	// 	{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10},
	// }))
	// [1,5], [1,10], [2,3], [5,10], [6,8]
}

// for i, _ := range intervals {
// 	s := startMap[i]
// 	e := startMap[i]
// 	delete(startMap, i)
// 	delete(endMap, i)
// 	for k, st := range startMap {
// 		end := endMap[k]
// 		if end < s || st > e {
// 			delete(startMap, k)
// 			delete(endMap, k)
// 		}
// 	}
// 	fmt.Println(startMap, endMap, count)
// 	count++
// }
