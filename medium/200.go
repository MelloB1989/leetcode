package main

import "fmt"

func numIslands(grid [][]byte) int {
	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	ans := 0
	for r, row := range grid {
		for c := range row {
			if grid[r][c] == '1' {
				ans++
				dfs(r, c)
			}
		}
	}
	return ans
}

func main() {
	// grid := [][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	numberOfIslands := numIslands(grid)
	fmt.Println(numberOfIslands)
}
