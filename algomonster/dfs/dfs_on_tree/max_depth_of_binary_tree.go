// Find longest root-to-leaf path

package main

import "fmt"

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	treeMaxDepth := func(root *Node) int {
		var dfs func(root *Node) int

		dfs = func(root *Node) int {
			if root == nil {
				return 0
			}

			leftDepth := dfs(root.left)
			rightDepth := dfs(root.right)

			if leftDepth >= rightDepth {
				return leftDepth + 1
			}

			return rightDepth + 1
		}

		// dfs() returns the number of nodes from root to leaf.
		// If we only want edges, subtract 1.
		if root != nil {
			return dfs(root) - 1
		}

		// Null tree
		return 0

	}

	tree := &Node{
		val: 5,
		left: &Node{
			val: 4,
			left: &Node{
				val: 3,
			},
			right: &Node{
				val: 8,
			},
		},
		right: &Node{
			val: 6,
		},
	}

	fmt.Println(treeMaxDepth(tree))

}
