package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	n := len(people)
	boats := 0
	l, r := 0, n-1
	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		boats++
	}
	return boats
}

func main() {
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
}
