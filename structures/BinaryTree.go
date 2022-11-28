package structures

import "golang.org/x/exp/constraints"

// TreeNode for int type only
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Binary Search Tree node for Ordered types
type BSTNode[T constraints.Ordered] struct {
	Val   T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

// Binary Search Tree (BST) Functions
// ------------------------------------------

// BST Search
// Returns true if found, false if not
func (n *BSTNode[T]) Search(key T) bool {
	if n == nil {
		return false
	}

	if key < n.Val {
		return n.Left.Search(key)
	} else if key > n.Val {
		return n.Right.Search(key)
	}

	return true
}

// BST Find max value
// TODO

// BST Find min value
// TODO

// Validate if BST
// TODO

// Binary Tree Functions
// ------------------------------------------

// Insert
func (n *TreeNode) Insert(val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else if val > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

// Delete

// Count nodes

// Calculate height

// Check existence

// Pre-order traversal: Root, Left, Right (recursive)
func preOrderTraverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preOrderTraverse(root.Left, result)
	preOrderTraverse(root.Right, result)
}
func PreOrderTraversal(root *TreeNode) []int {
	result := []int{}
	preOrderTraverse(root, &result)
	return result
}

// Pre-order traversal (iterative)
func PreOrderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		// Pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			result = append(result, node.Val)

			// Push if not nil
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return result
}

// In-order traversal: Left, Root, Right (recursive)
// Same as pre-order, but traverse left first.
func inOrderTraverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inOrderTraverse(root.Left, result)
	*result = append(*result, root.Val)
	inOrderTraverse(root.Right, result)
}
func InOrderTraversal(root *TreeNode) []int {
	result := []int{}
	inOrderTraverse(root, &result)
	return result
}

// In-order traversal (iterative)
func InOrderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) != 0 {
		// Traverse to the leftmost node, stacking each one
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Pop
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)

		current = current.Right
	}

	return result
}

// Post-order traversal: Left, Right, Root (recursive)
// Same as pre-order, but traverse left and right before appending root.
func postOrderTraverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	postOrderTraverse(root.Left, result)
	postOrderTraverse(root.Right, result)
	*result = append(*result, root.Val)
}
func PostOrderTraversal(root *TreeNode) []int {
	result := []int{}
	postOrderTraverse(root, &result)
	return result
}

// Post-order traversal (iterative)
func PostOrderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	var last *TreeNode

	for root != nil || len(stack) != 0 {
		if root != nil {
			// Traverse all the way leftward, pushing each node onto the stack.
			stack = append(stack, root)
			root = root.Left
		} else {
			// Peek at top of stack
			node := stack[len(stack)-1]

			// Go right if possible, unless we've already visited it.
			if node.Right != nil && last != node.Right {
				root = node.Right
			} else {
				result = append(result, node.Val)
				last = node
				stack = stack[:len(stack)-1]
			}
		}
	}
	return result
}

// Level-order traversal top-down
// TODO

// Level-order traversal bottom-up
// TODO
