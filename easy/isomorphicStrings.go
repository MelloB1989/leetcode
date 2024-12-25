package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var mapS, mapT [256]byte
	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]
		if mapS[sc] == 0 && mapT[tc] == 0 {
			mapS[sc] = tc
			mapT[tc] = sc
		} else {
			if mapS[sc] != tc || mapT[tc] != sc {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("add", "egg"))
	// fmt.Println(isIsomorphic("badc", "baba"))
}
