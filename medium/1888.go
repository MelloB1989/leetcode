package main

import (
	"fmt"
	"strconv"
)

func minFlips(s string) int {
	n := len(s)
	const INT_MAX = 1<<31 - 1
	ans1, ans2, ans := 0, 0, INT_MAX
	a := []int{}
	for i := 0; i < 2*n; i++ {
		if i < n {
			p, _ := strconv.Atoi(string(s[i]))
			a = append(a, p)
		}
		if i%2 != a[i%n] {
			ans1++
		}
		if (i+1)%2 != a[i%n] {
			ans2++
		}
		if i >= n {
			if (i-n)%2 != a[i-n] {
				ans1--
			}
			if (i-n+1)%2 != a[i-n] {
				ans2--
			}
		}
		if i >= n-1 {
			ans = min(min(ans1, ans2), ans)
		}
	}

	return ans
}

func main() {
	fmt.Println(minFlips("111000"))
}
