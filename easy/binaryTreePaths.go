package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	var b func(r *TreeNode)
	b = func(r *TreeNode) {
		if r.Left != nil {
			v := strconv.Itoa(r.Val)
			paths = append(paths, v)
			b(r.Left)
		} else if r.Right != nil {
			v := strconv.Itoa(r.Val)
			paths = append(paths, v)
			b(r.Right)
		}
		v := strconv.Itoa(r.Val)
		paths = append(paths, v)
		b(root.Right)
	}

	b(root)

	return paths
}

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}))
}
