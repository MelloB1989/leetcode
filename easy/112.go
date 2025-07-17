package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode, current_sum int) bool
	dfs = func(node *TreeNode, current_sum int) bool {
		if node == nil {
			return false
		}
		current_sum += node.Val
		if node.Right == nil && node.Left == nil && current_sum == targetSum {
			return true
		}
		return dfs(node.Left, current_sum) || dfs(node.Right, current_sum)
	}
	return dfs(root, 0)
}

func main() {
	fmt.Println(hasPathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}, 22))
}
