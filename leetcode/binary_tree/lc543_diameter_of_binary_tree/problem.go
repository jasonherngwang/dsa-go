/*
Diameter of Binary Tree
https://leetcode.com/problems/diameter-of-binary-tree/

Problem
------------------------------------------
Given the root of a binary tree, return the length of the diameter of the tree.

Inputs: Pointer to root node
Output: Diameter (integer)

Constraints:
The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100

Rules, Requirements, Definitions
------------------------------------------
"Diameter": Length of longest path between any two nodes.
- May or may not pass through the root.
- Path length is the number of edges between the nodes.

Assumptions
- The diameter is the sum of the depths of the left and right subtrees of a
  particular node.

Examples, Test Cases
------------------------------------------
Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1

Data Structure
------------------------------------------
Binary tree

Algorithm
------------------------------------------
Edge case: If root null return 0.

Approach 1: Recursive
For each node, recursively find the max depth of its two subtrees, and sum them
to yield the diameter. Keep track of the max diameter observed.

Time:
Space:

Steps
- Base case: If node doesn't exist, return 0
- Make recursive call to obtain the depth of both subtrees.
- Sum the two depths to yield the diameter.
  - If this is larger than the max diameter, update it.
- Return this node's max depth, plus 1 to account for the edge between it and
  its parent.
*/

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Approach 1: Recursive
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	leftDepth := dfs(root.Left, maxDiameter)
	rightDepth := dfs(root.Right, maxDiameter)

	diameter := leftDepth + rightDepth
	*maxDiameter = max(*maxDiameter, diameter)

	return max(leftDepth, rightDepth) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	dfs(root, &maxDiameter)
	return maxDiameter
}
