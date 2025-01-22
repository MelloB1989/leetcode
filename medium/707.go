package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	List   ListNode
	Head   *ListNode
	Tail   *ListNode
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		List:   ListNode{},
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	i := 0
	head := this.Head
	for head != nil {
		if index == i {
			return head.Val
		}
		head = head.Next
		i += 1
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: this.Head,
	}
	if this.Tail == nil {
		this.Tail = newNode
	}
	this.Head = newNode
	this.Length += 1
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	if this.Head == nil {
		this.Head = newNode
		this.Tail = newNode
	} else {
		this.Tail.Next = newNode
		this.Tail = newNode
	}
	this.Length += 1
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Length {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Length {
		this.AddAtTail(val)
		return
	}

	curr := this.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	newNode := &ListNode{
		Val:  val,
		Next: curr.Next,
	}
	curr.Next = newNode
	this.Length += 1
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Length {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	curr := this.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	this.Length -= 1
}

func (this *MyLinkedList) PrintList() {
	r := this.Head
	for r != nil {
		fmt.Print(r.Val, "->")
		r = r.Next
	}
	fmt.Println("nil")
}

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	fmt.Println(list.Get(0))
	list.AddAtIndex(1, 2)
	list.PrintList()
	list.DeleteAtIndex(1)
	list.PrintList()
}
