package main

import "fmt"

// func main() {
// 	a := []int{3, -5, 7, -9, -11, 1}
// 	b := make([]int, len(a))
// 	//[-5,-9,-11,1]
// 	i, j := 0, 1
// 	top := 0
// 	for i < len(a) && j < len(a) {
// 		if a[j] < 0 {
// 			if absInt(a[j]) > a[i] {
// 				b[top] = a[j]
// 				top += 1
// 			} else {
// 				b[top] = a[j]
// 				top += 1
// 			}
// 		}
// 		i = i + 1
// 		j = j + 1
// 	}
// 	fmt.Print(b)
// }

//	func absInt(x int) int {
//		if x < 0 {
//			return -x
//		}
//		return x
//	}
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, a := range asteroids {
		for len(stack) > 0 && a < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			fmt.Println(stack)
			if top < -a {
				stack = stack[:len(stack)-1]
				continue
			} else if top == -a {
				stack = stack[:len(stack)-1]
			}
			a = 0
			break
		}
		if a != 0 {
			stack = append(stack, a)
		}
	}

	return stack
}

func main() {
	a := []int{3, -5, 7, -9, -11, 1}
	result := asteroidCollision(a)
	fmt.Println(result) // Output will be [-5, -9, -11, 1]
}
