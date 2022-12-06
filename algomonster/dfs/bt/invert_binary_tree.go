/*
Invert Binary Tree
- In-place

At each node:
- Swap its children.
- Perform recursive DFS call both its children, to invert the subtrees.
*/
package main

import "fmt"

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	invertBinaryTree := func(tree *Node) *Node {
		var dfs func(node *Node) *Node
		dfs = func(node *Node) *Node {
			if node == nil {
				return node
			}
			// Invert current node
			node.left, node.right = node.right, node.left
			// Invert subtree
			dfs(node.left)
			dfs(node.right)

			return node
		}

		return dfs(tree)
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

	fmt.Println(invertBinaryTree(tree).left.left)
}
