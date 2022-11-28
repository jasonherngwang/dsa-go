package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertTree(t *testing.T) {
	t.Run("Insert numbers", func(t *testing.T) {
		tree := &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 10},
		}
		tree.Insert(3)
		tree.Insert(7)

		expected := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:  10,
				Left: &TreeNode{Val: 7},
			},
		}

		assert.Equal(t, expected, tree)
	})
}

func TestSearchBST(t *testing.T) {
	t.Run("Search for number", func(t *testing.T) {
		bst := &BSTNode[int]{
			Val: 5,
			Left: &BSTNode[int]{
				Val:   1,
				Right: &BSTNode[int]{Val: 3},
			},
			Right: &BSTNode[int]{
				Val:  10,
				Left: &BSTNode[int]{Val: 7},
			},
		}

		actual := true
		expected := bst.Search(7)

		assert.Equal(t, expected, actual)
	})
}
