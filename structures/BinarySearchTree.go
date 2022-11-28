package structures

import "golang.org/x/exp/constraints"

// TreeNode
type TreeNode[T constraints.Ordered] struct {
	Key   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Insert
// Base case is to do nothing
func (n *TreeNode[T]) Insert(key T) {
	if key < n.Key {
		// move left
		if n.Left == nil {
			n.Left = &TreeNode[T]{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		// move right
		if n.Right == nil {
			n.Right = &TreeNode[T]{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

// Search
// Returns true if found, false if not
func (n *TreeNode[T]) Search(key T) bool {
	if n == nil {
		return false
	}

	if key < n.Key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}

	return true
}
