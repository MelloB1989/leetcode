package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var cases int

	fmt.Fscan(in, &cases)

	for i := 0; i < cases; i++ {
		var n, x int
		fmt.Fscan(in, &n, &x)

		subs := math.Ceil(float64(n) / 6)
		fmt.Println(int(subs) * x)
	}
}
