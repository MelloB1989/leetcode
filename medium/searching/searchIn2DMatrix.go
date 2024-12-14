package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for rows := 0; rows < len(matrix); rows += 1 {
		first := 0
		last := len(matrix[0]) - 1
		if target <= matrix[rows][last] && target >= matrix[rows][first] {
			mid := (first + last) / 2
			for first <= last {
				mid = (first + last) / 2
				if matrix[rows][mid] == target {
					return true
				} else if matrix[rows][mid] < target {
					first = mid + 1
				} else {
					last = mid - 1
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 1))
}
