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
		var x int
		fmt.Fscan(in, &x)
		if x <= 7 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
