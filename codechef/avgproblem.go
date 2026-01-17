package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var cases int
	fmt.Fscan(in, &cases)

	for i := 0; i < cases; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		avg := float64(a+b) / 2.0
		if avg > float64(c) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
