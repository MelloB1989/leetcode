package main

import "fmt"

func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, len(prices))
	stack := []int{}
	copy(result, prices)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			result[j] = prices[j] - prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
