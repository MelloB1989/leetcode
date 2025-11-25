package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type KthLargest struct {
	k    int
	nums []int
	q    *priorityqueue.Queue
}

func Constructor(k int, nums []int) KthLargest {
	c := KthLargest{
		k:    k,
		nums: nums,
		q: priorityqueue.NewWith(func(a, b any) int {
			return cmp.Compare(a.(int), b.(int))
		}),
	}
	for _, n := range nums {
		c.q.Enqueue(n)
		if c.q.Size() > c.k {
			c.q.Dequeue()
		}
	}
	return c
}

func (this *KthLargest) Add(val int) int {
	this.q.Enqueue(val)
	if this.q.Size() > this.k {
		this.q.Dequeue()
	}
	res, _ := this.q.Peek()
	return res.(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
