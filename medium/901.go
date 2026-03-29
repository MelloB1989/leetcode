package main

import "fmt"

type Pair struct {
	price int
	span  int
}

type StockSpanner struct {
	stack []Pair
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []Pair{},
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stack) > 0 && this.stack[len(this.stack)-1].price <= price {
		span += this.stack[len(this.stack)-1].span
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, Pair{
		price: price,
		span:  span,
	})
	return span
}

func main() {
	sp := Constructor()
	for _, n := range []int{100, 80, 60, 70, 60, 75, 85} {
		fmt.Println(n, sp.Next(n))
	}
}
