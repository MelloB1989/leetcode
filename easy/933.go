package main

import (
	"fmt"
)

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.calls = append(this.calls, t)
	s := t - 3000
	for len(this.calls) > 0 && this.calls[0] < s {
		this.calls = this.calls[1:]
	}
	return len(this.calls)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {
	rc := Constructor()
	calls := []int{1, 100, 3001, 3002}
	ans := []int{}
	for _, c := range calls {
		ans = append(ans, rc.Ping(c))
	}
	fmt.Println(ans)
}
