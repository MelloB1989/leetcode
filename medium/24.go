package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new := &ListNode{Next: head}
	n := new
	for head != nil && head.Next != nil {
		first := head
		second := first.Next

		n.Next = second
		first.Next = second.Next
		second.Next = first

		n = first
		head = first.Next
	}
	return new.Next
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	r := swapPairs(list)
	for r != nil {
		println(r.Val)
		r = r.Next
	}
}
