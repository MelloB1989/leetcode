package main

import "fmt"

type MyCircularDeque struct {
	Deque    []int
	Front    int
	Rear     int
	Size     int
	Capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Deque:    make([]int, k),
		Front:    0,
		Rear:     0,
		Size:     0,
		Capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Deque[this.Front] = value
	} else {
		this.Front = (this.Front - 1 + this.Capacity) % this.Capacity
		this.Deque[this.Front] = value
	}
	this.Size++
	if this.Size == 1 {
		this.Rear = this.Front
	}
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Deque[this.Rear] = value
	} else {
		this.Rear = (this.Rear + 1) % this.Capacity
		this.Deque[this.Rear] = value
	}
	this.Size++
	if this.Size == 1 {
		this.Front = this.Rear
	}
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Size > 1 {
		this.Front = (this.Front + 1) % this.Capacity
	}
	this.Size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Size > 1 {
		this.Rear = (this.Rear - 1 + this.Capacity) % this.Capacity
	}
	this.Size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Deque[this.Front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Deque[this.Rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size == this.Capacity
}

func main() {
	myCircularDeque := Constructor(3)
	// fmt.Println(myCircularDeque.Rear, myCircularDeque.Front)
	fmt.Println(myCircularDeque.InsertLast(1)) // return True
	// fmt.Println(myCircularDeque.Rear, myCircularDeque.Front)
	fmt.Println(myCircularDeque.InsertLast(2))  // return True
	fmt.Println(myCircularDeque.InsertFront(3)) // return True
	fmt.Println(myCircularDeque.InsertFront(4)) // return False, the queue is full.
	// fmt.Println(myCircularDeque.Deque, myCircularDeque.Front, myCircularDeque.Rear)
	fmt.Println(myCircularDeque.GetRear())      // return 2
	fmt.Println(myCircularDeque.IsFull())       // return True
	fmt.Println(myCircularDeque.DeleteLast())   // return True
	fmt.Println(myCircularDeque.InsertFront(4)) // return True
	fmt.Println(myCircularDeque.GetFront())     // return 4
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
