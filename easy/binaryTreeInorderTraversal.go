package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	traverseInOrder(root, &result)
	return result
}

func traverseInOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	traverseInOrder(node.Left, result)
	*result = append(*result, node.Val)
	traverseInOrder(node.Right, result)
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Print(inorderTraversal(root))
}
