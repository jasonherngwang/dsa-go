/*
Balanced: Each node's subtrees have height diff <= 1

Info needed at each node to make a decision:
- Whether it has already been determined if one of its subtrees is imbalanced.
  - If imbalanced, -1 is returned.
  - If both its subtrees are balanced, compare their heights to check if the
    current subtree is balanced.
	- If imbalanced return -1
	- If balanced return the greater of its subtrees' height, plus 1 for the
	  current node.

Return value: Height of current tree, or -1 if imbalance detected.
Pass state down: None needed
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	var treeHeight func(root *Node) int

	treeHeight = func(root *Node) int {
		if root == nil {
			return 0
		}

		leftHeight := treeHeight(root.left)
		rightHeight := treeHeight(root.right)

		// If it has already been determined that a subtree is imbalanced
		if leftHeight == -1 || rightHeight == -1 {
			return -1
		}

		// Check if the current subtree is imbalanced
		if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
			return -1
		}

		// Otherwise return the current tree height
		if leftHeight >= rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	isBalanced := func(root *Node) bool {
		return treeHeight(root) != -1
	}

	tree := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 4,
				right: &Node{
					val: 7,
				},
			},
			right: &Node{
				val: 5,
			},
		},
		right: &Node{
			val: 3,
			right: &Node{
				val: 6,
			},
		},
	}
	tree2 := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 4,
				right: &Node{
					val: 7,
				},
			},
			right: &Node{
				val: 5,
			},
		},
		right: &Node{
			val: 3,
			right: &Node{
				val: 6,
				left: &Node{
					val: 8,
				},
			},
		},
	}

	fmt.Println(isBalanced(tree))
	fmt.Println(isBalanced(tree2))

}
