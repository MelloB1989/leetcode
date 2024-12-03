package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}}
	result := rotateBy90(arr)
	fmt.Println(result)
}

func rotateBy90(arr [][]int) [][]int {
	pr, pc := len(arr), len(arr[0])
	rows, cols := pc, pr
	array := make([][]int, rows)
	for i := range array {
		array[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			array[i][j] = arr[cols-j-1][i]
		}
	}
	return array
}
