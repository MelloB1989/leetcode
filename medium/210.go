package main

import "fmt"

type Queue struct {
	q      []int
	length int
}

func (q *Queue) push(v int) {
	q.q = append(q.q, v)
	q.length++
}
func (q *Queue) poll() int {
	if q.length == 0 {
		return -1
	}
	front := q.q[0]
	q.q = q.q[1:]
	q.length--
	return front
}
func findOrder(numCourses int, prerequisites [][]int) []int {
	reqMap := map[int][]int{}
	inDegree := make([]int, numCourses)
	for _, reqs := range prerequisites {
		reqMap[reqs[1]] = append(reqMap[reqs[1]], reqs[0])
		inDegree[reqs[0]]++
	}
	queue := Queue{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.push(i)
		}
	}
	result := []int{}
	for queue.length != 0 {
		curr := queue.poll()
		result = append(result, curr)

		for _, req := range reqMap[curr] {
			inDegree[req]--
			if inDegree[req] == 0 {
				queue.push(req)
			}
		}
	}
	if len(result) == numCourses {
		return result
	}
	return []int{}
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
