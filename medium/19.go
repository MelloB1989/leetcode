package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: head,
	}
	dist := res

	for _ = range n {
		head = head.Next
	}

	for head != nil {
		head = head.Next
		dist = dist.Next
	}
	dist.Next = dist.Next.Next

	return res.Next
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	r := removeNthFromEnd(list, 2)
	for r != nil {
		println(r.Val)
		r = r.Next
	}
}
