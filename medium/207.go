package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	reqMap := map[int][]int{}
	taken := map[int]bool{}
	for _, reqs := range prerequisites {
		reqMap[reqs[0]] = append(reqMap[reqs[0]], reqs[1])
	}
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if len(reqMap[course]) == 0 {
			return true
		}
		if taken[course] {
			return false
		}
		taken[course] = true
		for _, req := range reqMap[course] {
			if !dfs(req) {
				return false
			}
		}
		reqMap[course] = []int{}
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{0, 1}}))
}
