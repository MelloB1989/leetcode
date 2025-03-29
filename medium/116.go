package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	level := root
	for level != nil {
		current_node := level
		for current_node != nil {
			if current_node.Left != nil && current_node.Right != nil {
				current_node.Left.Next = current_node.Right
			}
			if current_node.Right != nil && current_node.Next != nil {
				current_node.Right.Next = current_node.Next.Left
			}
			current_node = current_node.Next
		}
		level = level.Left
	}
	return root
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("Node %d -> ", root.Val)
	if root.Next != nil {
		fmt.Printf("Next %d\n", root.Next.Val)
	} else {
		fmt.Println("Next nil")
	}
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}

	root2 := connect(root)
	printTree(root)
	printTree(root2)
}
