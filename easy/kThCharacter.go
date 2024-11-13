package main

import "fmt"

func kthCharacter(k int) byte {
	init := "a"
	for len(init) < k {
		init += genDo(init)
	}
	return init[k-1]
}

func genDo(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		c := int(s[i])
		r += string(c + 1)
		if c == 122 {
			r += string(97)
		}
	}
	return r
}

func main() {
	fmt.Println(kthCharacter(5))
}
