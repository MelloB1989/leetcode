package main

// abcabc
func numberOfSubstrings(s string) int {
	seen := map[rune]int{
		'a': -1,
		'b': -1,
		'c': -1,
	}
	count := 0
	for i, c := range s {
		seen[c] = i
		count += min(seen['a'], seen['b'], seen['c']) + 1
	}
	return count
}
