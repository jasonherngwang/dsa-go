/*
Invert Binary Tree
https://leetcode.com/problems/invert-binary-tree/

Problem
------------------------------------------
Given the root of a binary tree, invert the tree, and return its root.

Inputs: Pointer to root node

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Rules, Requirements, Definitions
------------------------------------------
"Invert" means mirroring a tree so that the left and right children are swapped.
- Root doesn't change.
- Imagine a vertical line running through Root, acting as the line to mirror
  across.

Examples, Test Cases
------------------------------------------
Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []

Data Structure
------------------------------------------
Binary tree

Algorithm
------------------------------------------
Approach 1: Recursive
Starting from root, recursively invoke an invert() function on its children.
This function swaps the children. Stop when the node has no children.
- If a node has only 1 child, we still perform the swap. Tree does not have to
  be complete.

Time: O(N) since we visit each node once
Space: O(N) recursive calls on the call stack

Approach 2: Iterative
- Starting from root, add both its children to a stack.
- While the stack is not empty:
  - Pop a node off the stack.
  - Swap its children.
  - Add its children to the stack.
*/

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Recursive
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// Approach 2: Iterative
func invertTreeIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	stack = append(stack, root)
	var curr *TreeNode

	for len(stack) != 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr != nil {
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil || curr.Right != nil {
				curr.Left, curr.Right = curr.Right, curr.Left
			}
		}
	}

	return root
}
