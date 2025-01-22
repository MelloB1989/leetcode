package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	current := head
	for current != nil {
		cur := current
		next := cur.Next
		temp := next
		next.Next = cur
		cur.Next = temp.Next
		current = current.Next
	}
	return head
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	r := swapPairs(list)
	for r != nil {
		println(r.Val)
		r = r.Next
	}
}
