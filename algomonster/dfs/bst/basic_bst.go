package main

import "fmt"

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	var search func(node *Node, val int) bool
	search = func(node *Node, val int) bool {
		if node == nil {
			return false
		}

		if val < node.val {
			return search(node.left, val)
		} else if val > node.val {
			return search(node.right, val)
		}

		return true
	}

	// Insert
	var insert func(node *Node, val int) *Node
	insert = func(node *Node, val int) *Node {
		if node == nil {
			return &Node{val: val}
		}

		if val < node.val {
			node.left = insert(node.left, val)
		} else if val > node.val {
			node.right = insert(node.right, val)
		}

		return node
	}

	tree := &Node{
		val: 5,
		left: &Node{
			val:   1,
			right: &Node{val: 3},
		},
		right: &Node{
			val:   10,
			left:  &Node{val: 7},
			right: &Node{val: 12},
		},
	}

	fmt.Println(search(tree, 7))
	fmt.Println(search(tree, 42))

	insert(tree, 2)
	insert(tree, 8)
	fmt.Println(search(tree, 2))
	fmt.Println(search(tree, 8))
}
