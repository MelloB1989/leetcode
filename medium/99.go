package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var dfs func(node *TreeNode, min, max *int)
	dfs = func(node *TreeNode, min, max *int) {
		if node == nil {
			return
		}
		if max != nil && node.Val > *max {
			temp := *min
			*min = node.Val
			node.Val = temp
		}
		if min != nil && node.Val < *min {
			temp := *max
			*max = node.Val
			node.Val = temp
		}
		dfs(node.Left, min, &node.Val)
		dfs(node.Right, &node.Val, max)
	}
}

func main() {
	recoverTree(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
	})
}
