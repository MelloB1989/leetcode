package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	MOD := 1_000_000_007
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)

	stack := []int{}

	for i, num := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = 1 + i
		} else {
			left[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n - i
		} else {
			right[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	sum := 0
	for i, v := range arr {
		sum = (sum + v*left[i]*right[i]) % MOD
	}
	return sum
}

func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}
