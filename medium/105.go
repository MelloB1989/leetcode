package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inIndexMap := map[int]int{}
	for i, n := range inorder {
		inIndexMap[n] = i
	}

	var build func(preStart, inStart, inEnd int) *TreeNode

	build = func(preStart, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		val := preorder[preStart]
		inIndex := inIndexMap[val]
		root := &TreeNode{
			Val: val,
		}
		root.Left = build(preStart+1, inStart, inIndex-1)
		root.Right = build(preStart+(inIndex-inStart)+1, inIndex+1, inEnd)
		return root
	}

	return build(0, 0, len(inorder)-1)
}
