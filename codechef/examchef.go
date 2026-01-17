package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var c int
	fmt.Fscan(in, &c)

	for i := 0; i < c; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		stds := x * y
		per := float64(z) / float64(stds)
		if per > 0.5 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
