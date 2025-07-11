package main

import "fmt"

func maxProfit(k int, prices []int) int {
	mins := make([]int, k)
	for i := range mins {
		mins[i] = prices[0]
	}

	txs := make([]int, k)

	for _, price := range prices {
		mins[0] = min(mins[0], price)
		txs[0] = max(txs[0], price-mins[0])
		for i := 1; i < k; i++ {
			mins[i] = min(mins[i], price-txs[i-1])
			txs[i] = max(txs[i], price-mins[i])
		}
	}

	return txs[k-1]
}

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
}
