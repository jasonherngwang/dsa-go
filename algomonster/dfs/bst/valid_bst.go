/*
Valid BST

In-order traversal (left, current, right) yields a monotonically increasing list
of values

Return value required to make decision at node: Boolean indicating whether
node's subtree is a valid BST.
- Computed by performing DFS on both its children.
- Null nodes are valid BSTs, so return true if null.

Pass state down? Upper and lower bounds so the child can check its value against
them.
Left child:
- Lower: -Infinity (min int)
- Upper: Parent's value
Right child:
- Lower: Parent's value
- Upper: Infinity (max int)

Steps
- Perform in-order traversal:
  - If nil return true.
  - If out-of-range, return false
  - Return recursive DFS on both children.
- Initiate DFS using -Inf and Inf as range bounds.

Time: O(N) to visit all node
Space: O(log N) for tree height, worst case O(N) for linked-list
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

	validBst := func(root *Node) bool {
		var dfs func(root *Node, min int, max int) bool
		dfs = func(root *Node, min int, max int) bool {
			// Ignore empty nodes
			if root == nil {
				return true
			}

			// Node value is out-of-range; not a valid BST
			if !(root.val > min && root.val < max) {
				return false
			}

			// Continue checking
			return dfs(root.left, min, root.val) && dfs(root.right, root.val, max)
		}

		return dfs(root, math.MinInt, math.MaxInt)
	}

	valid := &Node{
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

	fmt.Println(validBst(valid))

	invalid := &Node{
		val: 5,
		left: &Node{
			val:   1,
			right: &Node{val: 3},
		},
		right: &Node{
			val:   10,
			left:  &Node{val: 7},
			right: &Node{val: 9}, // Should be right child of 7
		},
	}

	fmt.Println(validBst(invalid))
}
