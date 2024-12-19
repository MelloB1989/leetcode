package main

import "fmt"

func distinctNames(ideas []string) int64 {
	count := int64(0)
	ideaMap := make(map[rune][]string)

	getDuplicates := func(a []string, b []string) int {
		m := map[string]int{}
		d := 0
		for _, v := range a {
			m[v] += 1
		}
		for _, v := range b {
			m[v] += 1
		}
		for _, v := range m {
			if v > 1 {
				d += 1
			}
		}
		return d
	}

	for _, idea := range ideas {
		ideaMap[rune(idea[0])] = append(ideaMap[rune(idea[0])], idea[1:])
	}

	for i := 'a'; i <= 'z'; i++ {
		for j := i + 1; j <= 'z'; j++ {
			if len(ideaMap[i]) > 0 && len(ideaMap[j]) > 0 {
				duplicates := getDuplicates(ideaMap[i], ideaMap[j])
				// fmt.Println(ideaMap[i], ideaMap[j], duplicates)
				count += int64(2 * (len(ideaMap[i]) - duplicates) * (len(ideaMap[j]) - duplicates))
			}
		}
	}

	return count
}

func main() {
	// fmt.Println(distinctNames([]string{"lack", "back"}))
	fmt.Println(distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
}
