package main

import "fmt"

func suggestedProducts(products []string, searchWord string) [][]string {
	results := [][]string{}
	//Sort lexicographically
	for i := 0; i < len(products); i++ {
		for j := i + 1; j < len(products); j++ {
			if products[i] > products[j] {
				products[i], products[j] = products[j], products[i]
			}
		}
	}
	getWords := func(key string) []string {
		n := len(key)
		r := []string{}
		for _, word := range products {
			if len(word) < n {
				continue
			}
			sub := word[:n]
			if sub == key {
				r = append(r, word)
			}
		}
		if len(r) < 3 {
			return r
		}
		return r[:3]
	}
	for i := 1; i <= len(searchWord); i++ {
		results = append(results, getWords(searchWord[:i]))
	}
	return results
}

func main() {
	fmt.Println(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"))
	// fmt.Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
}
