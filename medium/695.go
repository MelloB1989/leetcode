package main

func maxAreaOfIsland(grid [][]int) int {
	var getArea func(i, j int) int
	getArea = func(i, j int) int {
		if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0

		return 1 +
			getArea(i+1, j) +
			getArea(i-1, j) +
			getArea(i, j+1) +
			getArea(i, j-1)
	}
	maxArea := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 1 {
				maxArea = max(getArea(i, j), maxArea)
			}
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
