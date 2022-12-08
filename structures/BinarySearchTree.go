package structures

import "golang.org/x/exp/constraints"

// Binary Search Tree node for Ordered types
type BSTNode[T constraints.Ordered] struct {
	Val   T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

// Search: Returns true if found, false if not
func (node *BSTNode[T]) Search(val T) bool {
	if node == nil {
		return false
	}

	if val < node.Val {
		return node.Left.Search(val)
	} else if val > node.Val {
		return node.Right.Search(val)
	}

	return true
}

// Insert
func (node *BSTNode[T]) Insert(val T) *BSTNode[T] {
	if node == nil {
		return &BSTNode[T]{Val: val}
	}

	if val < node.Val {
		node.Left = node.Left.Insert(val)
	} else if val > node.Val {
		node.Right = node.Right.Insert(val)
	}

	return node
}

// Find max value
// TODO

// Find min value
// TODO

// Validate if BST
// TODO
