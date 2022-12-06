package main

import "fmt"

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	insertBst := func(bst *Node, val int) *Node {
		var dfs func(node *Node, val int) *Node
		dfs = func(node *Node, val int) *Node {
			if node == nil {
				return &Node{val: val}
			}
			if val < node.val {
				node.left = dfs(node.left, val)
			} else if val > node.val {
				node.right = dfs(node.right, val)
			}

			return node
		}
		return dfs(bst, val)
	}

	tree := &Node{
		val: 5,
		left: &Node{
			val:   1,
			right: &Node{val: 3},
		},
		right: &Node{
			val:  10,
			left: &Node{val: 7},
			// right: &Node{val: 12},
		},
	}

	insertBst(tree, 12)
	fmt.Println(tree.right.right)
}
