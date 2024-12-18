package main

import "fmt"

func minGroups(intervals [][]int) int {
	startMap := make(map[int]int)
	endMap := make(map[int]int)
	count := 0

	for i, interval := range intervals {
		startMap[i] = interval[0]
		endMap[i] = interval[1]
	}

	for len(startMap) > 0 {
		var firstKey int
		for k := range startMap {
			firstKey = k
			break
		}
		fmt.Println(len(startMap), intervals[firstKey])

		s := startMap[firstKey]
		e := endMap[firstKey]
		delete(startMap, firstKey)
		delete(endMap, firstKey)

		toDelete := []int{}
		for k, st := range startMap {
			end := endMap[k]
			if end < s || st > e {
				toDelete = append(toDelete, k)
				if st < s {
					s = st
				}
				if end > e {
					e = end
				}
			}
		}

		for _, k := range toDelete {
			delete(startMap, k)
			delete(endMap, k)
		}

		count++
	}

	return count
}

func main() {
	fmt.Println(minGroups([][]int{
		{1, 3}, {5, 6}, {8, 10}, {11, 13},
	}))
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
