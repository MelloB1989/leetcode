package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	prev := head
	current := head
	for current != nil {
		if current.Val == prev.Val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}
	return head
}

func main() {
	sortedList := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	finalList := deleteDuplicates(sortedList)
	//Print the final list
	node := finalList
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
}
