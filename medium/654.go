package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := []*TreeNode{}
	for _, n := range nums {
		node := &TreeNode{Val: n}
		var last *TreeNode
		for len(stack) > 0 && stack[len(stack)-1] != nil && stack[len(stack)-1].Val < n {
			last = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		node.Left = last
		if len(stack) != 0 {
			stack[len(stack)-1].Right = node
		}
		stack = append(stack, node)
	}
	if len(stack) == 0 {
		return nil
	}
	return stack[0]
}

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	printTree(root, "", true)
}

// helper function
func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	if node.Right != nil {
		newPrefix := prefix
		if isLeft {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		printTree(node.Right, newPrefix, false)
	}

	if isLeft {
		fmt.Println(prefix + "└── " + fmt.Sprint(node.Val))
	} else {
		fmt.Println(prefix + "┌── " + fmt.Sprint(node.Val))
	}

	if node.Left != nil {
		newPrefix := prefix
		if isLeft {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		printTree(node.Left, newPrefix, true)
	}
}
