package main

import (
	"fmt"
)

// func shoppingOffers(price []int, special [][]int, needs []int) int {
// 	//Choose a offer to go with
// 	var chooseBestOffer func(n []int) []int
// 	chooseBestOffer = func(n []int) []int {
// 		best := -1
// 		off := math.MaxInt
// 		for o, offer := range special {
// 			offerValid := true
// 			of := offer[len(offer)-1]
// 			for i, need := range n {
// 				if !(offer[i] <= need) {
// 					offerValid = false
// 					break
// 				}
// 			}
// 			if offerValid && of < off {
// 				off = of
// 				best = o
// 			}
// 		}
// 		if best == -1 {
// 			return []int{}
// 		}
// 		return special[best]
// 	}

// 	var applyOffer func(offer []int, n []int) ([]int, int, int)
// 	applyOffer = func(offer []int, n []int) ([]int, int, int) { //needs, spend, itemsLeft
// 		newNeeds := make([]int, len(n))
// 		copy(newNeeds, n)
// 		leftItems := 0
// 		for i := range newNeeds {
// 			newNeeds[i] -= offer[i]
// 			leftItems += newNeeds[i]
// 		}
// 		return newNeeds, offer[len(offer)-1], leftItems
// 	}

// 	var buyActual func(n []int) ([]int, int, int) //needs, spend, itemsLeft
// 	buyActual = func(n []int) ([]int, int, int) {
// 		TS := 0
// 		newNeeds := make([]int, len(n))
// 		copy(newNeeds, n)
// 		for i, p := range price {
// 			TS += p * newNeeds[i]
// 			newNeeds[i] = 0
// 		}
// 		return newNeeds, TS, 0
// 	}

// 	totalSpend := 0
// 	leftItems := 0
// 	currentNeeds := make([]int, len(needs))
// 	copy(currentNeeds, needs)

// 	for _, items := range currentNeeds {
// 		leftItems += items
// 	}

// 	for leftItems > 0 {
// 		off := chooseBestOffer(currentNeeds)
// 		// fmt.Println(off)
// 		_, actual, _ := buyActual(currentNeeds)
// 		if len(off) != 0 {
// 			n, spend, l := applyOffer(off, currentNeeds)
// 			// fmt.Println(needs, spend, leftItems)
// 			if spend < actual {
// 				currentNeeds = n
// 				leftItems = l
// 				totalSpend += spend
// 			} else {
// 				currentNeeds, totalSpend, leftItems = buyActual(currentNeeds)
// 			}
// 		} else {
// 			for i, p := range price {
// 				totalSpend += p * currentNeeds[i]
// 				leftItems -= currentNeeds[i]
// 				currentNeeds[i] = 0
// 			}
// 			// fmt.Println(needs)
// 		}
// 	}

//		return totalSpend
//	}
func shoppingOffers(price []int, special [][]int, needs []int) int {
	var calculateRegularPrice func(n []int) int
	calculateRegularPrice = func(n []int) int {
		total := 0
		for i, p := range price {
			total += p * n[i]
		}
		return total
	}

	// DFS with memoization
	memo := make(map[string]int)

	var dfs func(currentNeeds []int) int
	dfs = func(currentNeeds []int) int {
		// Convert needs to string for memoization
		key := fmt.Sprint(currentNeeds)
		if val, exists := memo[key]; exists {
			return val
		}

		// Calculate regular price
		minCost := calculateRegularPrice(currentNeeds)

		// Try each special offer
		for _, offer := range special {
			valid := true
			remainingNeeds := make([]int, len(currentNeeds))
			copy(remainingNeeds, currentNeeds)

			// Check if offer can be applied
			for i := 0; i < len(currentNeeds); i++ {
				remainingNeeds[i] -= offer[i]
				if remainingNeeds[i] < 0 {
					valid = false
					break
				}
			}

			if valid {
				// Cost of current offer + minimum cost for remaining needs
				currentCost := offer[len(offer)-1] + dfs(remainingNeeds)
				minCost = min(minCost, currentCost)
			}
		}

		memo[key] = minCost
		return minCost
	}

	return dfs(needs)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	price := []int{2, 3, 4}
	special := [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}
	needs := []int{1, 2, 1}
	// price := []int{0, 0, 0}
	// special := [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}
	// needs := []int{1, 1, 1}
	fmt.Println(shoppingOffers(price, special, needs))
}
