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
	var b func(r *TreeNode, path string)
	b = func(r *TreeNode, path string) {
		if r == nil {
			return
		}
		if len(path) > 0 {
			path += "->"
		}
		path += strconv.Itoa(r.Val)
		if r.Left == nil && r.Right == nil {
			paths = append(paths, path)
		} else {
			b(r.Left, path)
			b(r.Right, path)
		}
	}

	b(root, "")

	return paths
}

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}))
}
