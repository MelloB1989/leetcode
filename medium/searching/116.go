package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levels := root
	for levels != nil {
		current_node := levels
		for current_node != nil {
			if current_node.Left != nil {
				fmt.Println("Connecting ", current_node.Left, current_node.Right)
				current_node.Left.Next = current_node.Right
			}
			if current_node.Right != nil && current_node.Next != nil {
				fmt.Println("Connecting ", current_node.Right, current_node.Next.Left)
				current_node.Right.Next = current_node.Next.Left
			}
			current_node = current_node.Next
		}
		levels = levels.Left
	}
	return root
}

func main() {
	r := connect(
		&Node{
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
			},
		},
	)
	level := r
	for level != nil {
		node := level
		for node != nil {
			fmt.Print(node.Val, " -> ")
			node = node.Next
		}
		fmt.Println("NULL")
		level = level.Left
	}
}
