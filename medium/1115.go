package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type FooBar struct {
	n int
	f chan bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n, f: make(chan bool)}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.f <- true
		<-fb.f
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.f
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.f <- true
	}
}

func main() {
	fb := NewFooBar(2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fb.Foo(func() {
			fmt.Print("foo")
		})
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fb.Bar(func() {
			fmt.Print("bar")
		})
	}()
	wg.Wait()
}
