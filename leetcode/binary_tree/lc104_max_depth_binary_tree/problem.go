/*
Invert Binary Tree
https://leetcode.com/problems/maximum-depth-of-binary-tree/

Problem
------------------------------------------
Given the root of a binary tree, return its maximum depth.

Inputs: Pointer to root node
Output: Max depth (integer)

Constraints:
The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100

Rules, Requirements, Definitions
------------------------------------------
"Max depth": Number of nodes (not edges) along the longest path from root down
to the farthest leaf node.

Examples, Test Cases
------------------------------------------
Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3 (3->20-15 or 3->20-7)

Example 2:
Input: root = [1,null,2]
Output: 2

Data Structure
------------------------------------------
Binary tree

Algorithm
------------------------------------------
Edge case: If root null return 0.

Approach 1: Recursive
Recursively traverse:
- Base case: If node doesn't exist, return 0
- If leaf node, return current depth + 1
- If either child exists, make recursive call to obtain the depth of that
subtree.
  - Return the deeper of the two subtrees, plus 1.

Time: O(N) since we visit each node once
Space: O(N) recursive calls on the call stack

Approach 2: Iterative (using queue)
- Initialize counter variable to 0.
- Initialize queue (as slice) and enqueue root.
  - Queue front is index 0. Queue end is last index.
- While queue is not empty:
  - Measure queue length. We only want to increment by 1 for each level of the
    tree, not for each node.
  - Iterate from 0 to queue length - 1:
    - Dequeue from queue, and add its children to the queue.
  - Increment depth by 1.
*/

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Recursive
func traverse(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}

	// Leaf
	if root.Left == nil && root.Right == nil {
		return depth + 1
	}

	// Has child
	leftHeight := traverse(root.Left, depth)
	rightHeight := traverse(root.Right, depth)

	if leftHeight >= rightHeight {
		return depth + leftHeight + 1
	}
	return depth + rightHeight + 1
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 0)
}

// Approach 2: Iterative
func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0

	queue := []*TreeNode{}
	queue = append(queue, root)

	var node *TreeNode

	for len(queue) > 0 {
		treeLevelSize := len(queue)

		for i := 0; i < treeLevelSize; i++ {
			node = queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}
