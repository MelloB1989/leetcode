package main

import (
	"fmt"
	"log"
)

func solveNQueens(n int) [][]string {
	result := [][]string{}
	currentPath := []string{}

	var b func(posx int)
	b = func(posx int) {
		log.Println(currentPath)
		if len(currentPath) == n {
			s := make([]string, n)
			copy(s, currentPath)
			result = append(result, s)
			return
		}
		for j := 1; j <= n; j++ {
			if isPlacable(j, currentPath) {
				currentPath = append(currentPath, genPosString(j, n))
				b(posx + 1)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
	}
	b(0)
	return result
}

func isPlacable(cur int, currentPath []string) bool {
	curRow := len(currentPath)
	for row, col := range currentPath {
		p := getPosInt(col)
		if p == cur || abs(p-cur) == abs(row-curRow) {
			return false
		}
	}
	return true
}

func genPosString(posy, n int) string {
	r := ""
	for i := 1; i <= n; i++ {
		if i == posy {
			r += "Q"
		} else {
			r += "."
		}
	}
	return r
}

func getPosInt(posy string) int {
	for i, s := range posy {
		if s == 'Q' {
			return i + 1
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(solveNQueens(4))
}
