package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	new := &ListNode{Next: head}

	for head == nil || head.Next == nil {
		return head
	}
	length := 1
	t := head
	for t.Next != nil {
		length += 1
		t = t.Next
	}

	k = k % length
	if k == 0 {
		return head
	}
	cur := head
	for i := 0; i < length-k-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	t.Next = new.Next
	// fmt.Println(cur, t)
	return newHead
}

func main() {
	// list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list := &ListNode{1, &ListNode{2, nil}}
	r := rotateRight(list, 1)
	for r != nil {
		println(r.Val)
		r = r.Next
	}
}
