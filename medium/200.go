package main

import "fmt"

func numIslands(grid [][]byte) int {
	x, y := 0, 0
	rows, cols := len(grid), len(grid[0])
	for x < rows && y < cols {

	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	// grid1 := [][]byte{
	// 	{'1', '1', '0', '0', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '1', '0', '0'},
	// 	{'0', '0', '0', '1', '1'},
	// }

	numberOfIslands := numIslands(grid)
	fmt.Println(numberOfIslands)
}
