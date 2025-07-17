package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode, current_sum int, leaves []int)
	dfs = func(node *TreeNode, current_sum int, leaves []int) {
		if node == nil {
			return
		}
		current_sum += node.Val
		leaves = append(leaves, node.Val)
		if node.Left == nil && node.Right == nil && current_sum == targetSum {
			path := make([]int, len(leaves))
			copy(path, leaves)
			result = append(result, path)
		}
		dfs(node.Left, current_sum, leaves)
		dfs(node.Right, current_sum, leaves)
	}
	dfs(root, 0, []int{})
	return result
}

func main() {
	fmt.Println(pathSum(&TreeNode{
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
				Left:  &TreeNode{Val: 5},
			},
		},
	}, 22))
}
