package main

import (
"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func PreorderRecursive(root *Node) {
	if root != nil {
		fmt.Printf("%d , ", root.val)
		PreorderRecursive(root.left)
		PreorderRecursive(root.right)
	}
}
func PostorderRecursive(root *Node) {
	if root != nil {

		PostorderRecursive(root.left)
		PostorderRecursive(root.right)
		fmt.Printf("%d , ", root.val)
	}
}

func InorderRecursive(root *Node) {
	if root != nil {
		InorderRecursive(root.left)
		fmt.Printf("%d , ", root.val)
		InorderRecursive(root.right)
	}
}

func TreeTraversInPrePost() {
/*
			10
		   /  \
		 20	   30
		/ \      \
	   40  50     60
	  /
	 70
*/

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Preorder Traversal - Recursive Solution : ")
	PreorderRecursive(root)
	fmt.Println("\nPostorder Traversal - Recursive Solution : ")
	PostorderRecursive(root)
	fmt.Println("\nInorder Traversal - Recursive Solution : ")
	InorderRecursive(root)
	fmt.Println()
}