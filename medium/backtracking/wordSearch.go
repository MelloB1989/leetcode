package main

import "fmt"

func exist(board [][]byte, word string) bool {
	var dfs func(index int, r int, c int) bool
	rows := len(board)
	cols := len(board[0])
	dfs = func(index int, r int, c int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		found := dfs(index+1, r-1, c) || //Up
			dfs(index+1, r+1, c) || //Down
			dfs(index+1, r, c-1) || //Left
			dfs(index+1, r, c+1) //Right

		board[r][c] = temp
		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}

	return false
}

func main() {
	//fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABAB"))
	fmt.Println(exist([][]byte{{'a', 'b'}, {'c', 'd'}}, "acdb"))
}
