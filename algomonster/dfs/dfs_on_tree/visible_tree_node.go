/*
Visible means it is the greatest on its path to root.

Output: Number of visible nodes

Recursive return value: Count of visible nodes in subtree

States passed down: Largest value between current node and root.
- This is used for comparison with current node.

Notes
- If a node < parent, there could still be visible nodes in it descendants.
- "Greatest" is greater or equal

Algo
Base case: Null node; return 0

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

	visibleTreeNode := func(root *Node) int {
		var dfs func(node *Node, max int) int
		dfs = func(node *Node, max int) int {
			if node == nil {
				return 0
			}

			count := 0
			if node.val >= max {
				max = node.val
				count++
			}

			count += dfs(node.left, max)
			count += dfs(node.right, max)

			return count
		}

		return dfs(root, math.MinInt)
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

	fmt.Println(visibleTreeNode(tree))

}
