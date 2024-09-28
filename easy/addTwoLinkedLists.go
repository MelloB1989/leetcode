package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var resultHead *ListNode = &ListNode{Val: 0, Next: nil}
	_ = recursiveAddition(l1, l2, &carry, resultHead)
	return resultHead
}

func recursiveAddition(l1 *ListNode, l2 *ListNode, carry *int, head *ListNode) *ListNode {
	if l1 == nil && l2 == nil && *carry == 0 {
		return nil
	}
	digit1 := 0
	digit2 := 0
	var next1 *ListNode = nil
	var next2 *ListNode = nil
	if l1 != nil {
		digit1 = l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		digit2 = l2.Val
		next2 = l2.Next
	}
	sum := digit1 + digit2 + *carry
	//fmt.Print(digit1, "+", digit2, "+", *carry, "=", sum%10, "\n")
	head.Val = sum % 10
	*carry = sum / 10
	if next1 != nil || next2 != nil || *carry != 0 {
		temp := &ListNode{Val: 0, Next: nil}
		head.Next = temp
		recursiveAddition(next1, next2, carry, temp)
	}
	return head
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

func main() {
	list1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	list2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result := addTwoNumbers(list1, list2)
	printList(result)
}
