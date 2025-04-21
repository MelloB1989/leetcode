// Input: s = "aabacbebebe", k = 3
// Output: 7
// Explanation: "cbebebe" is the longest substring with 3 distinct characters.
// Input: s = "aaaa", k = 2
// Output: -1
// Explanation: There's no substring with 2 distinct characters.
// Input: s = "aabaaab", k = 2
// Output: 7
// Explanation: "aabaaab" is the longest substring with 2 distinct characters.

package main

import "fmt"

func longestKUniqueCharactersSubstring(s string, k int) int {
	max, l := -1, 0
	m := map[byte]int{}
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		if len(m) > k {
			m[s[l]]--
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		} else if len(m) == k && len(s[l:r+1]) > max {
			max = len(s[l : r+1])
		}
	}
	return max
}

func main() {
	s := "aabacbebebe"
	k := 3
	fmt.Println(longestKUniqueCharactersSubstring(s, k))
}
