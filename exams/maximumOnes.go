package main

func maximumOnes(matrix [][]int) int {
	result := -1
	maxOnes := -1
	for r, row := range matrix {
		count := 0
		for j := len(row) - 1; j >= 0; j-- {
			count += row[j]
			if row[j] == 0 {
				break
			}
		}
		if count > maxOnes {
			maxOnes = count
			result = r
		}
	}
	return result
}
