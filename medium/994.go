package main

import "fmt"

type Queue struct {
	q      []any
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v any) {
	q.q = append(q.q, v)
	q.length++
}

func (q *Queue) Poll() any {
	if q.length == 0 {
		return -1
	}
	front := q.q[0]
	q.q = q.q[1:]
	q.length--
	return front
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Flush() {
	q.q = []any{}
	q.length = 0
}

func orangesRotting(grid [][]int) int {
	type coordinates struct {
		x int
		y int
	}
	rows := len(grid)
	cols := len(grid[0])
	visited := map[coordinates]bool{}
	queue := NewQueue()
	freshCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue.Push(coordinates{x: i, y: j})
				visited[coordinates{x: i, y: j}] = true
			}
			if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	directions := []coordinates{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	mins := -1
	for !queue.IsEmpty() {
		lvlSize := queue.length
		for i := 0; i < lvlSize; i++ {
			curr := queue.Poll().(coordinates)
			for _, dir := range directions {
				newCord := coordinates{
					x: curr.x + dir.x,
					y: curr.y + dir.y,
				}
				if newCord.x > -1 && newCord.x < rows && newCord.y > -1 && newCord.y < cols && grid[newCord.x][newCord.y] == 1 && !visited[newCord] {
					grid[newCord.x][newCord.y] = 2
					visited[newCord] = true
					queue.Push(newCord)
					freshCount--
				}
			}
		}
		mins++
	}
	if freshCount == 0 {
		return mins
	}
	return -1
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}
