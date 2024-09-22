package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	sub := strs[0]
	str := ""
	flag := 0
	for i := 0; i < len(strs); i++ {
		if len(sub) > len(strs[i]) {
			sub = sub[0:len(strs[i])]
			str = strs[i]
		} else {
			str = (string(strs[i]))[0:len(sub)]
		}
		for j := len(sub); j >= 0; j-- {
			if str[0:j] == sub[0:j] {
				sub = str[0:j]
				flag = 1
				break
			}
		}
	}
	if flag == 1 {
		return sub
	} else {
		return ""
	}
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
