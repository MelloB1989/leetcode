package main

import "fmt"

type Charset map[byte]bool

func (c Charset) add(char byte) {
	c[char] = true
}

func (c Charset) contains(char byte) bool {
	if _, ok := c[char]; !ok {
		return false
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	rows := map[int]Charset{}
	cols := map[int]Charset{}
	boxes := map[int]Charset{}
	for i := 0; i < 9; i++ {
		rows[i] = make(Charset)
		cols[i] = make(Charset)
		boxes[i] = make(Charset)
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {

			if board[r][c] == '.' {
				continue
			}

			// Check rows
			if ok := rows[r].contains(board[r][c]); ok {
				return false
			}
			rows[r].add(board[r][c])

			//Check cols
			if ok := cols[c].contains(board[r][c]); ok {
				return false
			}
			cols[c].add(board[r][c])

			//Check boxes
			boxIndex := (r/3)*3 + (c / 3)
			if ok := boxes[boxIndex].contains(board[r][c]); ok {
				return false
			}
			boxes[boxIndex].add(board[r][c])
		}
	}

	return true
}

func main() {
	fmt.Println(isValidSudoku([][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}
