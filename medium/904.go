package main

import "fmt"

func totalFruit(fruits []int) int {
	n := len(fruits)
	baskets := map[int]int{}
	maxFruits := 0
	if n == 0 {
		return 0
	}

	left := 0
	for right := 0; right < n; right++ {
		baskets[fruits[right]]++
		for len(baskets) > 2 {
			baskets[fruits[left]]--
			if baskets[fruits[left]] == 0 {
				delete(baskets, fruits[left])
			}
			left++
		}
		collectedFruits := right - left + 1
		maxFruits = max(collectedFruits, maxFruits)
	}

	return maxFruits
}

func main() {
	// fmt.Println(totalFruit([]int{1, 0, 3, 4, 3}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
}
