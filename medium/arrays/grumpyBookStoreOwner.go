package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	maxUnsatisfied := 0
	currentUnsatisfied := 0
	totalSatisfied := 0

	for i, c := range customers[:minutes] {
		currentUnsatisfied += c * grumpy[i]
		totalSatisfied += c * (1 - grumpy[i])
	}
	maxUnsatisfied = currentUnsatisfied

	for i := minutes; i < len(customers); i++ {
		currentUnsatisfied += customers[i] * grumpy[i]
		currentUnsatisfied -= customers[i-minutes] * grumpy[i-minutes]
		maxUnsatisfied = max(currentUnsatisfied, maxUnsatisfied)
		totalSatisfied += customers[i] * (1 - grumpy[i])
	}

	return totalSatisfied + maxUnsatisfied
}

func main() {
	customers := []int{4, 10, 10}
	grumpy := []int{1, 1, 0}
	minutes := 2
	// customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	// grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	// minutes := 3
	result := maxSatisfied(customers, grumpy, minutes)
	fmt.Println(result)
}
