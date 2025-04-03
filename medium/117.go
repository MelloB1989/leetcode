package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize-1; i++ {
			queue[i].Next = queue[i+1]
		}
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
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
	// root := &Node{
	// 	Val: 1,
	// 	Left: &Node{
	// 		Val: 2,
	// 		Left: &Node{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 			Next:  nil,
	// 		},
	// 		Right: nil,
	// 		Next:  nil,
	// 	},
	// 	Right: &Node{
	// 		Val:  3,
	// 		Left: nil,
	// 		Right: &Node{
	// 			Val:   5,
	// 			Left:  nil,
	// 			Right: nil,
	// 			Next:  nil,
	// 		},
	// 		Next: nil,
	// 	},
	// 	Next: nil,
	// }
	// root := &Node{
	// 	Val: 1,
	// 	Left: &Node{
	// 		Val: 2,
	// 		Left: &Node{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 			Next:  nil,
	// 		},
	// 		Right: &Node{
	// 			Val:   5,
	// 			Left:  nil,
	// 			Right: nil,
	// 			Next:  nil,
	// 		},
	// 		Next: nil,
	// 	},
	// 	Right: &Node{
	// 		Val:  3,
	// 		Left: nil,
	// 		Right: &Node{
	// 			Val:   7,
	// 			Left:  nil,
	// 			Right: nil,
	// 			Next:  nil,
	// 		},
	// 		Next: nil,
	// 	},
	// 	Next: nil,
	// }

	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
				Left: &Node{
					Val:   7,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
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
			Val:  3,
			Left: nil,
			Right: &Node{
				Val:  6,
				Left: nil,
				Right: &Node{
					Val:   8,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Next: nil,
		},
		Next: nil,
	}

	fmt.Println("Before:")
	printTree(root)

	root = connect(root)
	fmt.Println("After:")
	printTree(root)
}
