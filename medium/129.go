package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, path int)
	sum := 0
	dfs = func(node *TreeNode, path int) {
		if node == nil {
			return
		}
		path = path*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sum += path
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, 0)
	return sum
}
