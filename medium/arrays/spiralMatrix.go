package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	result := spiralMatrix(arr)
	fmt.Println(result)
}

func spiralMatrix(arr [][]int) []int {
	rows, cols := len(arr), len(arr[0])
	result := make([]int, 0, rows*cols)
	top, bottom, left, right := 0, rows-1, 0, cols-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, arr[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, arr[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, arr[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, arr[i][left])
			}
			left++
		}
	}
	return result
}
