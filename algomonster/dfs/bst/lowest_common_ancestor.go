/*
Lowest Common Ancestor of BST
LCA is the "deepest" ancestor node of two given nodes p and q.
- Assume unique values; p and q appear once each in the tree.
- There may be nodes higher up (e.g. root) that are a common ancestor, but they
  are not the LCA.
- p and q can be descendants of themselves.
- For BST, left child < node, and right child > node.
  - If p and q are smaller than current node, discard right subtree and search
    left subtree only. Vice versa.
  - If current node is equal to either p or q, then the other value must be its
    descendant. Return itself.

The LCA of this would be p:
      p

q           5

Return value: Boolean indicating whether current node's subtree has p or q.

Pass state: p and q

Steps
- DFS. At each node:
  - Check if p and q are smaller than current. If so, return result of DFS on
    left subtree.
  - Check if p and q are larger than current. If so, return result of DFS on
    right subtree.
  - Return current value.

Time: O(log N) since we are splitting the search space each time, not
traversing all N nodes.
Space: O(h) = O(log N) for call stack

Worst case O(N) for both (skewed tree)
*/

package main

import "fmt"

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	var lcaOnBst func(bst *Node, p int, q int) int
	lcaOnBst = func(bst *Node, p int, q int) int {
		if p < bst.val && q < bst.val {
			return lcaOnBst(bst.left, p, q)
		} else if p > bst.val && q > bst.val {
			return lcaOnBst(bst.right, p, q)
		} else {
			return bst.val
		}
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
	fmt.Println(lcaOnBst(tree, 7, 12)) // 10
	fmt.Println(lcaOnBst(tree, 7, 3))  // 5
}
