package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	new := &ListNode{Next: head}
	n := new

	for _ = range k {
		head = head.Next
	}

	for head.Next != nil {
		n = n.Next
		head = head.Next
	}

	f := n.Next.Next
	n.Next.Next = nil
	head.Next = new.Next
	// fmt.Println(n.Next, head, f)

	return f
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	r := rotateRight(list, 2)
	for r != nil {
		println(r.Val)
		r = r.Next
	}
}
