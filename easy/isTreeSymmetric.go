package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	fmt.Println(isSymmetric(root))
}
