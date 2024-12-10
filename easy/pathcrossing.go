package main

import "fmt"

func isPathCrossing(path string) bool {
	myJourney := make(map[[2]int]bool)
	x, y := 0, 0

	myJourney[[2]int{x, y}] = true

	for _, direction := range path {
		switch direction {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		if myJourney[[2]int{x, y}] {
			return true
		}

		myJourney[[2]int{x, y}] = true
	}
	return false
}

func main() {
	fmt.Println(isPathCrossing("NEEEESWNN"))
}
