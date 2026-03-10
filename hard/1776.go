package main

import "fmt"

type stack struct {
	d []int
	l int
}

func createStack() *stack {
	return &stack{
		d: []int{},
		l: 0,
	}
}

func (s *stack) push(x int) {
	s.d = append(s.d, x)
	s.l++
}

func (s *stack) pop() int {
	if s.l == 0 {
		return -1
	}
	x := s.d[s.l-1]
	s.l--
	s.d = s.d[:s.l]
	return x
}

func (s *stack) getTop() int {
	if s.l == 0 {
		return -1
	}
	return s.d[s.l-1]
}

func (s *stack) isEmpty() bool {
	return s.l == 0
}

func getCollisionTimes(cars [][]int) []float64 {
	n := len(cars)
	stack := createStack()
	result := make([]float64, n)
	for i := range result {
		result[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		for !stack.isEmpty() {
			j := stack.getTop()

			if cars[i][1] > cars[j][1] {
				t := float64(cars[j][0]-cars[i][0]) / float64(cars[i][1]-cars[j][1])
				if result[j] != -1 && result[j] < t {
					stack.pop()
				} else {
					result[i] = t
					break
				}
			} else {
				stack.pop()
			}
		}

		stack.push(i)
	}

	return result
}

func main() {
	fmt.Println(getCollisionTimes([][]int{
		[]int{1, 2},
		[]int{2, 1},
		[]int{4, 3},
		[]int{7, 2},
	}))
}
