package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type H2O struct {
	h    chan struct{}
	o    chan struct{}
	cond *sync.Cond
}

func NewH2O() *H2O {
	h := &H2O{
		h:    make(chan struct{}, 2),
		o:    make(chan struct{}, 1),
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return h
}

func (h *H2O) releaseH2O() {
	if len(h.h) == 2 && len(h.o) == 1 {
		<-h.h
		<-h.h
		<-h.o
		h.cond.Broadcast()
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.cond.L.Lock()
	for len(h.h) == 2 {
		h.cond.Wait()
	}
	h.h <- struct{}{}
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	h.releaseH2O()
	h.cond.L.Unlock()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.cond.L.Lock()
	for len(h.o) == 1 {
		h.cond.Wait()
	}
	h.o <- struct{}{}
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	h.releaseH2O()
	h.cond.L.Unlock()
}

func main() {
	h2o := NewH2O()

	input := []rune("HHHHOOHHO")

	wg := sync.WaitGroup{}

	wg.Add(len(input))
	for _, c := range input {
		if c == 'H' {
			go func() {
				defer wg.Done()
				h2o.Hydrogen(func() {
					fmt.Print("H")
				})
			}()
		} else {
			go func() {
				defer wg.Done()
				h2o.Oxygen(func() {
					fmt.Print("O")
				})
			}()
		}
	}
	wg.Wait()
}
