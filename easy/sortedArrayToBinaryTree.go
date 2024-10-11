package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		nums[len(nums)/2],
		sortedArrayToBST(nums[:len(nums)/2]),
		sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

func PrintTreeGraphically(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	// Prepare the prefix for child nodes
	newPrefix := prefix
	if isLeft {
		fmt.Printf("%s├── %d\n", prefix, node.Val)
		newPrefix += "│   "
	} else {
		fmt.Printf("%s└── %d\n", prefix, node.Val)
		newPrefix += "    "
	}

	// Recursively print the left and right children
	PrintTreeGraphically(node.Left, newPrefix, true)
	PrintTreeGraphically(node.Right, newPrefix, false)
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	PrintTreeGraphically(root, "", false)
}
