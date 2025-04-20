package main

import "fmt"

func calculateScore(instructions []string, values []int) int64 {
	visited := map[int]bool{}
	var score int64
	i := 0
	for i < len(instructions) {
		if visited[i] {
			return score
		}
		visited[i] = true
		switch instructions[i] {
		case "jump":
			i += values[i]
			if i > len(instructions) || i < 0 {
				return score
			}
		case "add":
			score += int64(values[i])
			i++
		}
	}
	return score
}

func main() {
	instructions := []string{"jump", "add", "add"}
	values := []int{3, 1, 1}
	fmt.Println(calculateScore(instructions, values))
}
