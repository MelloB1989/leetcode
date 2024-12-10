package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	left := maxWeight(weights)
	right := sumWeights(weights)

	var canShip func(c int) bool
	canShip = func(c int) bool {
		currentLoad := 0
		reqDays := 1

		for _, w := range weights {
			if currentLoad+w > c {
				reqDays += 1
				currentLoad = 0
			}
			currentLoad += w

			if reqDays > days {
				return false
			}
		}

		return true
	}

	for left < right {
		mid := (left + right) / 2
		if canShip(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func maxWeight(w []int) int {
	max := w[0]
	for _, ww := range w {
		if max < ww {
			max = ww
		}
	}
	return max
}

func sumWeights(w []int) int {
	sum := 0
	for _, ww := range w {
		sum += ww
	}
	return sum
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
