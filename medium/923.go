package main

import (
	"fmt"
	"sort"
)

func threeSumMulti(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)

	const MOD = 1_000_000_007
	result := 0

	for i := 0; i < n; i++ {
		T := target - arr[i]

		l, r := i+1, n-1
		for l < r {
			sum := arr[l] + arr[r]
			if sum < T {
				l++
			} else if sum > T {
				r--
			} else {
				if arr[l] != arr[r] {
					lcount, rcount := 1, 1
					for l+1 < r && arr[l] == arr[l+1] {
						lcount++
						l++
					}
					for l < r-1 && arr[r] == arr[r-1] {
						rcount++
						r--
					}
					result = (result + lcount*rcount) % MOD
					r--
					l++
				} else {
					count := r - l + 1
					result = (result + count*(count-1)/2) % MOD
					break
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}
