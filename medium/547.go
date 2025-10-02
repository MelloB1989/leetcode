package main

import "fmt"

type queue struct{ a []int }

func newQueue() *queue      { return &queue{} }
func (q *queue) push(x int) { q.a = append(q.a, x) }
func (q *queue) pop() int {
	x := q.a[0]
	q.a = q.a[1:]
	return x
}
func (q *queue) isEmpty() bool { return len(q.a) == 0 }

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	ans := 0
	q := newQueue()

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		ans++
		q.push(i)
		visited[i] = true
		for !q.isEmpty() {
			cur := q.pop()
			for j, val := range isConnected[cur] {
				if val == 1 && !visited[j] {
					visited[j] = true
					q.push(j)
				}
			}
		}
	}

	return ans
}

func main() {
	isConnnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(isConnnected))
}
